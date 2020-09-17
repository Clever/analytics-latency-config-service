package config

import (
	"fmt"
	"testing"
	"time"

	"github.com/Clever/analytics-latency-config-service/gen-go/models"
	"github.com/Clever/analytics-latency-config-service/helpers"

	"github.com/stretchr/testify/assert"
)

func setup() {
	latencyConfigPath = "latency_config.json"
}

// TestParseChecks verifies that the latency_config
// JSON can successfully be unmarshalled.
// If this test fails, it is likely that there is a
// formatting issue with the latency_config
func TestParseChecks(t *testing.T) {
	setup()
	assert.NotPanics(t, func() {
		ParseChecks()
	}, "Unable to parse latency checks")
}

func TestValidateConfig(t *testing.T) {
	setup()
	configs := ParseChecks()
	dbs := []models.AnalyticsDatabase{
		models.AnalyticsDatabaseRedshiftProd,
		models.AnalyticsDatabaseRedshiftFast,
		models.AnalyticsDatabaseRdsInternal,
		models.AnalyticsDatabaseRdsExternal,
	}

	for _, db := range dbs {
		dbConfig, err := helpers.GetDatabaseConfig(configs, db)
		assert.Nilf(t, err, "invalid db config %s", db)
		assert.NotNilf(t, dbConfig, "db config for %s should not be nil", db)

		for _, schema := range dbConfig {
			assert.NotEmpty(t, schema.SchemaName, "schema config must specify a name")
			name := schema.SchemaName
			assert.NotEmptyf(t, schema.DefaultTimestampColumn, "schema %s config must specify a timestamp column", name)

			if len(schema.Whitelist) != 0 && len(schema.Blacklist) != 0 {
				assert.Failf(t, "schema %s can only contain a whitelist or a blacklist, not both.", name)
			}

			validateThresholds(t, schema.DefaultThresholds, name)

			for _, table := range schema.TableOverrides {
				assert.NotEmptyf(t, table.TableName, "table overrides for %s must specify a name", name)
				name := fmt.Sprintf("%s.%s", name, table.TableName)
				if table.LatencySpec != nil {
					validateThresholds(t, table.LatencySpec.Thresholds, name)
				}
			}
		}
	}
}

func validateThresholds(t *testing.T, thresholds *models.Thresholds, name string) {
	if thresholds == nil {
		return
	}

	tiers := []models.ThresholdTier{
		models.ThresholdTierCritical,
		models.ThresholdTierMajor,
		models.ThresholdTierMinor,
	}

	for _, tier := range tiers {
		threshold, err := helpers.GetThresholdTierValue(thresholds, tier)
		assert.Nilf(t, err, "invalid threshold tier: %s in %s", tier, name)

		if threshold == helpers.NoLatencyAlert || threshold == "" {
			continue
		}

		_, err = time.ParseDuration(threshold)
		assert.Nilf(t, err, "invalid duration %s for tier: %s in %s", threshold, tier, name)
	}
}
