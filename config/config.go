package config

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
	"path"

	"github.com/Clever/analytics-latency-config-service/gen-go/models"
	"github.com/Clever/analytics-latency-config-service/helpers"

	"github.com/kardianos/osext"
)

var (
	// We have two redshift databases:
	// One that holds all the data and views (prod)
	// And one that holds timeline (fast-prod)
	// RedshiftProd* are for the former
	RedshiftProdHost     string
	RedshiftProdPort     string
	RedshiftProdDatabase string
	RedshiftProdUsername string
	RedshiftProdPassword string

	// RedshiftFast* are for the latter
	RedshiftFastHost     string
	RedshiftFastPort     string
	RedshiftFastDatabase string
	RedshiftFastUsername string
	RedshiftFastPassword string

	// We also have two postgres Amazon RDS databases.
	// One that's for internal use (e..g building blocks)
	// And one that's for external use (e.g. district analytics.)
	// RDSInternal* is the former
	RDSInternalHost     string
	RDSInternalPort     string
	RDSInternalDatabase string
	RDSInternalUsername string
	RDSInternalPassword string

	// RDSExternal* is the former
	RDSExternalHost     string
	RDSExternalPort     string
	RDSExternalDatabase string
	RDSExternalUsername string
	RDSExternalPassword string

	DefaultOwner            string
	globalDefaultLatency    string
	GlobalDefaultThresholds models.Thresholds

	latencyConfigPath string
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

	dir, err := osext.ExecutableFolder()
	if err != nil {
		log.Fatal(err)
	}
	latencyConfigPath = path.Join(dir, "config/latency_config.json")
}

// Init reads environment variables and initializes the config.
func Init() {
	RedshiftProdHost = requiredEnv("REDSHIFT_PROD_HOST")
	RedshiftProdPort = requiredEnv("REDSHIFT_PROD_PORT")
	RedshiftProdDatabase = requiredEnv("REDSHIFT_PROD_DATABASE")
	RedshiftProdUsername = requiredEnv("REDSHIFT_PROD_USER")
	RedshiftProdPassword = requiredEnv("REDSHIFT_PROD_PASSWORD")

	RedshiftFastHost = requiredEnv("REDSHIFT_FAST_HOST")
	RedshiftFastPort = requiredEnv("REDSHIFT_FAST_PORT")
	RedshiftFastDatabase = requiredEnv("REDSHIFT_FAST_DATABASE")
	RedshiftFastUsername = requiredEnv("REDSHIFT_FAST_USER")
	RedshiftFastPassword = requiredEnv("REDSHIFT_FAST_PASSWORD")

	RDSInternalHost = requiredEnv("RDS_INTERNAL_HOST")
	RDSInternalPort = requiredEnv("RDS_INTERNAL_PORT")
	RDSInternalDatabase = requiredEnv("RDS_INTERNAL_DATABASE")
	RDSInternalUsername = requiredEnv("RDS_INTERNAL_USER")
	RDSInternalPassword = requiredEnv("RDS_INTERNAL_PASSWORD")

	RDSExternalHost = requiredEnv("RDS_EXTERNAL_HOST")
	RDSExternalPort = requiredEnv("RDS_EXTERNAL_PORT")
	RDSExternalDatabase = requiredEnv("RDS_EXTERNAL_DATABASE")
	RDSExternalUsername = requiredEnv("RDS_EXTERNAL_USER")
	RDSExternalPassword = requiredEnv("RDS_EXTERNAL_PASSWORD")
}

// ParseChecks reads in the latency check definitions
func ParseChecks() models.AnalyticsLatencyConfigs {
	latencyJSON, err := ioutil.ReadFile(latencyConfigPath)
	if err != nil {
		log.Fatalf("read-latency-config-error: %s", err.Error())
		panic("Unable to read latency config")
	}

	var checks models.AnalyticsLatencyConfigs
	err = json.Unmarshal(latencyJSON, &checks)
	if err != nil {
		log.Fatalf("parse-latency-checks-error: %s", err.Error())
		panic("Unable to parse latency checks")
	}

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
