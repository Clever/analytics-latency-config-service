package db

import (
	"fmt"
	"os"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func setup(t *testing.T) *Postgres {
	conf := PostgresCredentials{
		Host:     os.Getenv("POSTGRES_HOST"),
		Port:     os.Getenv("POSTGRES_PORT"),
		Username: os.Getenv("POSTGRES_USER"),
		Password: "",
		Database: os.Getenv("POSTGRES_DB"),
	}
	db, err := NewPostgresClient(conf, "testCluster")
	assert.NoError(t, err)

	_, err = db.GetSession().Exec("CREATE SCHEMA IF NOT EXISTS test")
	assert.NoError(t, err)

	_, err = db.GetSession().Exec("DROP TABLE IF EXISTS latencies")
	assert.NoError(t, err)

	_, err = db.GetSession().Exec(`CREATE TABLE latencies(name varchar(255), last_update timestamp without time zone);`)
	assert.NoError(t, err)

	return db
}

func TestQueryLatencyTableCheck(t *testing.T) {
	n := time.Now().In(time.UTC)
	past := time.Date(n.Year(), n.Month(), n.Day(), n.Hour()-96, 0, 0, 0, time.UTC)

	db := setup(t)

	latency, valid, err := db.QueryLatencyTable("test", "latency")
	assert.NoError(t, err)
	assert.False(t, valid)

	_, err = db.GetSession().Exec(fmt.Sprintf("INSERT INTO latencies(name, last_update) VALUES ('test.latency', '%s')",
		past.In(time.UTC).Format(time.RFC3339)))
	require.NoError(t, err)

	latency, valid, err = db.QueryLatencyTable("test", "latency")
	assert.NoError(t, err)
	assert.True(t, valid)
	// Give a little leeway for timing
	assert.True(t, latency >= 95 && latency <= 97)
}
