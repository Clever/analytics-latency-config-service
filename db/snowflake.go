package db

import (
	"database/sql"
	"fmt"
	"time"

	"github.com/Clever/analytics-latency-config-service/config"

	_ "github.com/snowflakedb/gosnowflake" // Snowflake driver
)

// SnowflakeCredentials provides an implementation of DBClient
// that contains the postgres client connection.
type Snowflake struct {
	session *sql.DB
}

// SnowflakeCredentials contains the Snowflake credentials/information.
type SnowflakeCredentials struct {
	Username      string
	Password      string
	Account       string
	Database      string
	Role          string
	Warehouse     string
	Authenticator string
}

// NewPostgresClient creates a Postgres db client.
func NewSnowflakeClient(info SnowflakeCredentials) (*Snowflake, error) {
	connectionParams := fmt.Sprintf("%s:%s@%s/%s?role=%s&warehouse=%s",
		info.Username, info.Password, info.Account, info.Database, info.Role, info.Warehouse)
	// Connection for local development
	if info.Authenticator != "" {
		connectionParams = fmt.Sprintf("%s@%s/%s?role=%s&warehouse=%s&authenticator=%s",
			info.Username, info.Account, info.Database, info.Role, info.Warehouse, info.Authenticator)
	}
	session, err := sql.Open("snowflake", connectionParams)

	if err != nil {
		return nil, err
	}

	return &Snowflake{session}, nil
}

// NewProdSnowflakeClient initializes a client to Snowflake
func NewSnowflakeProdClient() (*Snowflake, error) {
	info := SnowflakeCredentials{
		Account:       config.SnowflakeAccount,
		Username:      config.SnowflakeUsername,
		Password:      config.SnowflakePassword,
		Database:      config.SnowflakeDatabase,
		Warehouse:     config.SnowflakeWarehouse,
		Role:          config.SnowflakeRole,
		Authenticator: config.SnowflakeAuthenticator,
	}

	return NewSnowflakeClient(info)
}

// GetClusterName returns "snowflake" (more relevant for postgres/redshift)
func (c *Snowflake) GetClusterName() string {
	return "snowflake"
}

// GetSession returns the Snowflake session
func (c *Snowflake) GetSession() *sql.DB {
	return c.session
}

func (c *Snowflake) QueryLatencyTable(schemaName, tableName string) (int64, bool, error) {
	check := fmt.Sprintf("'%s.%s'", schemaName, tableName)
	if schemaName == "PUBLIC" {
		// We don't always prepend the schema name if it's "public". Check both.
		check = fmt.Sprintf("%s OR name = '%s'", check, tableName)
	}

	latencyQuery := fmt.Sprintf("SELECT extract(epoch from last_update) FROM latencies WHERE name = %s", check)
	var latency sql.NullFloat64
	err := c.GetSession().QueryRow(latencyQuery).Scan(&latency)
	if err != nil {
		if err == sql.ErrNoRows {
			return 0, false, nil
		}
		return 0, false, fmt.Errorf("error scanning latency table for %s: %s", check, err)
	}
	if !latency.Valid {
		return 0, false, nil
	}

	hourDiff := (time.Now().Unix() - int64(latency.Float64)) / 3600
	return hourDiff, latency.Valid, nil
}

// QueryTableMetadata returns a map of tables
// belonging to a given schema in Postgres, indexed
// by table name.
// It also attempts to infer the timestamp column, by
// choosing the alphabetically lowest column with a
// timestamp type. We use this as a heuristic since a
// lot of our timestamp columns are prefixed with "_".
func (c *Snowflake) QueryTableMetadata(schemaName string) (map[string]TableMetadata, error) {
	query := fmt.Sprintf(`
		SELECT table_name, min("column_name")
		FROM information_schema.columns
		WHERE table_schema ILIKE '%s'
		AND data_type ILIKE '%%timestamp%%'
		GROUP BY table_name
	`, schemaName)

	tableMetadata := make(map[string]TableMetadata)
	rows, err := c.GetSession().Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var row TableMetadata
		if err := rows.Scan(&row.TableName, &row.TimestampColumn); err != nil {
			return tableMetadata, fmt.Errorf("Unable to scan row for schema %s: %s", schemaName, err)
		}

		tableMetadata[row.TableName] = row
	}

	return tableMetadata, nil
}
