package db

import (
	"database/sql"
	"fmt"
	"log"
	"time"

	"github.com/Clever/analytics-latency-config-service/config"

	// Use our own version of the Postgres library so we get keep-alive support
	// See: https://github.com/Clever/pq/pull/1
	_ "github.com/Clever/pq"
)

// PostgresClient provides a default implementation of PostgresClient
// that contains the postgres client connection.
type Postgres struct {
	session     *sql.DB
	clusterName string
}

type PostgresClient interface {
	GetClusterName() string
	GetSession() *sql.DB
	QueryLatencyTable(schemaName, tableName string) (int64, bool, error)
}

// PostgresCredentials contains the postgres credentials/information.
type PostgresCredentials struct {
	Host     string
	Port     string
	Username string
	Password string
	Database string
}

// NewPostgresClient creates a Postgres db client.
func NewPostgresClient(info PostgresCredentials, clusterName string) (*Postgres, error) {
	const connectionTimeout = 60
	connectionParams := fmt.Sprintf("host=%s port=%s dbname=%s keepalive=1 connect_timeout=%d",
		info.Host, info.Port, info.Database, connectionTimeout)
	credentialsParams := fmt.Sprintf("user=%s password=%s", info.Username, info.Password)
	if info.Host == "localhost" {
		fmt.Println("nossl")
		// Locally we have to disable ssl mode
		connectionParams += " sslmode=disable"
	}

	log.Printf("new-postgres-client: connectionParams %+v", connectionParams)
	openParams := fmt.Sprintf("%s %s", connectionParams, credentialsParams)
	session, err := sql.Open("postgres", openParams)
	if err != nil {
		return nil, err
	}

	return &Postgres{session, clusterName}, nil
}

// NewRedshiftProdClient initializes a client to fresh prod
func NewRedshiftProdClient() (*Postgres, error) {
	info := PostgresCredentials{
		Host:     config.RedshiftProdHost,
		Port:     config.RedshiftProdPort,
		Username: config.RedshiftProdUsername,
		Password: config.RedshiftProdPassword,
		Database: config.RedshiftProdDatabase,
	}

	return NewPostgresClient(info, "redshift-prod")
}

// NewRedshiftFastClient initializes a client to fast prod
func NewRedshiftFastClient() (*Postgres, error) {
	info := PostgresCredentials{
		Host:     config.RedshiftFastHost,
		Port:     config.RedshiftFastPort,
		Username: config.RedshiftFastUsername,
		Password: config.RedshiftFastPassword,
		Database: config.RedshiftFastDatabase,
	}

	return NewPostgresClient(info, "redshift-fast")
}

// NewRDSInternalClient initializes a client to internal rds
func NewRDSInternalClient() (*Postgres, error) {
	info := PostgresCredentials{
		Host:     config.RDSInternalHost,
		Port:     config.RDSInternalPort,
		Username: config.RDSInternalUsername,
		Password: config.RDSInternalPassword,
		Database: config.RDSInternalDatabase,
	}

	return NewPostgresClient(info, "rds-internal")
}

// NewRDSExternalClient initializes a client to external rds
func NewRDSExternalClient() (*Postgres, error) {
	info := PostgresCredentials{
		Host:     config.RDSExternalHost,
		Port:     config.RDSExternalPort,
		Username: config.RDSExternalUsername,
		Password: config.RDSExternalPassword,
		Database: config.RDSExternalDatabase,
	}

	return NewPostgresClient(info, "rds-external")
}

// GetClusterName returns the name of the client Postgres cluster
func (c *Postgres) GetClusterName() string {
	return c.clusterName
}

// GetSession returns the name of the client Postgres cluster
func (c *Postgres) GetSession() *sql.DB {
	return c.session
}

func (c *Postgres) QueryLatencyTable(schemaName, tableName string) (int64, bool, error) {
	check := fmt.Sprintf("'%s.%s'", schemaName, tableName)
	if schemaName == "public" {
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
