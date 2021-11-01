package db

import "database/sql"

type DBClient interface {
	GetClusterName() string
	GetSession() *sql.DB
	QueryTableMetadata(schemaName string) (map[string]TableMetadata, error)
	QueryLatencyTable(schemaName, tableName string) (int64, bool, error)
}

// TableMetadata contains information about a table in Postgres
type TableMetadata struct {
	TableName       string
	TimestampColumn string
}
