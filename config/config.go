package config

import (
	"log"
	"os"

	"github.com/ghodss/yaml"

	"github.com/Clever/analytics-latency-config-service/gen-go/models"
	"github.com/Clever/analytics-latency-config-service/helpers"
)

var (
	// RedshiftFast* holds timeline (fast-prod)
	RedshiftFastHost     string
	RedshiftFastPort     string
	RedshiftFastDatabase string
	RedshiftFastUsername string
	RedshiftFastPassword string

	// We also have a postgres Amazon RDS databases for external use (e.g. district analytics.)

	// RDSExternal* is the former
	RDSExternalHost     string
	RDSExternalPort     string
	RDSExternalDatabase string
	RDSExternalUsername string
	RDSExternalPassword string

	// We have one Snowflake instance
	SnowflakeUsername      string
	SnowflakePassword      string
	SnowflakeAccount       string
	SnowflakeDatabase      string
	SnowflakeWarehouse     string
	SnowflakeRole          string
	SnowflakeAuthenticator string

	DefaultOwner            string
	globalDefaultLatency    string
	GlobalDefaultThresholds models.Thresholds

	latencyConfig string
)

func init() {
	DefaultOwner = "eng-deip"
	globalDefaultLatency = "24h"
	GlobalDefaultThresholds = models.Thresholds{
		Critical: globalDefaultLatency,
		Major:    globalDefaultLatency,
		Minor:    globalDefaultLatency,
		Refresh:  helpers.NoLatencyAlert,
	}
}

func Init() {
	InitDBs()
	InitConfig()
}

// InitDBs reads environment DB variables and initializes the config.
func InitDBs() {
	RedshiftFastHost = requiredEnv("REDSHIFT_FAST_HOST")
	RedshiftFastPort = requiredEnv("REDSHIFT_FAST_PORT")
	RedshiftFastDatabase = requiredEnv("REDSHIFT_FAST_DATABASE")
	RedshiftFastUsername = requiredEnv("REDSHIFT_FAST_USER")
	RedshiftFastPassword = requiredEnv("REDSHIFT_FAST_PASSWORD")

	RDSExternalHost = requiredEnv("RDS_EXTERNAL_HOST")
	RDSExternalPort = requiredEnv("RDS_EXTERNAL_PORT")
	RDSExternalDatabase = requiredEnv("RDS_EXTERNAL_DATABASE")
	RDSExternalUsername = requiredEnv("RDS_EXTERNAL_USER")
	RDSExternalPassword = requiredEnv("RDS_EXTERNAL_PASSWORD")

	SnowflakeUsername = requiredEnv("SNOWFLAKE_USER")
	SnowflakePassword = requiredEnv("SNOWFLAKE_PASSWORD")
	SnowflakeAccount = requiredEnv("SNOWFLAKE_ACCOUNT")
	SnowflakeDatabase = requiredEnv("SNOWFLAKE_DATABASE")
	SnowflakeWarehouse = requiredEnv("SNOWFLAKE_WAREHOUSE")
	SnowflakeRole = requiredEnv("SNOWFLAKE_ROLE")
	SnowflakeAuthenticator = os.Getenv("SNOWFLAKE_AUTHENTICATOR") // for local development
}

// InitConfig reads environment latency config variables and initializes the config.
// We separate this from the dbs so that we can call it separately.
func InitConfig() {
	latencyConfig = requiredEnv("LATENCY_CONFIG")
}

// ParseChecks reads in the latency check definitions
func ParseChecks() models.AnalyticsLatencyConfigs {
	if latencyConfig == "" {
		log.Fatalf("empty latency config")
		panic("Unable to read latency config")
	}

	var checks models.AnalyticsLatencyConfigs
	err := yaml.Unmarshal([]byte(latencyConfig), &checks)
	if err != nil {
		log.Fatalf("parse-latency-checks-error: %s", err.Error())
		panic("Unable to parse latency checks")
	}

	// Now that latency-config is an environment variable, we'll do validation
	// at startup. This way, if there's something wrong with the config, the service
	// will fail fast (and rollback) rather than have strange, intermittent errors.
	validateLatencyConfig(checks)

	return checks
}

// requiredEnv tries to find a value in the environment variables. If a value is not
// found the program will panic.
func requiredEnv(key string) string {
	value := os.Getenv(key)
	if value == "" {
		log.Fatalf("required-env: name: %s", key)
		os.Exit(1)
	}
	return value
}
