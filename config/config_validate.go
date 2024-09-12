package config

import (
	"fmt"
	"time"

	"github.com/Clever/analytics-latency-config-service/gen-go/models"
	"github.com/Clever/analytics-latency-config-service/helpers"
)

func validateLatencyConfig(configs models.AnalyticsLatencyConfigs) {
	dbs := []models.AnalyticsDatabase{
		models.AnalyticsDatabaseRedshiftFast,
		models.AnalyticsDatabaseRdsExternal,
		models.AnalyticsDatabaseSnowflake,
	}

	for _, db := range dbs {
		dbConfig, err := helpers.GetDatabaseConfig(configs, db)
		assertNilf(err, "invalid db config %s", db)
		assertNotNilf(dbConfig, "db config for %s should not be nil", db)

		for _, schema := range dbConfig {
			assertNotEmpty(schema.SchemaName, "schema config must specify a name")
			name := schema.SchemaName
			assertNotEmptyf(schema.DefaultTimestampColumn, "schema %s config must specify a timestamp column", name)

			if len(schema.Whitelist) != 0 && len(schema.Blacklist) != 0 {
				assertFailf("schema %s can only contain a whitelist or a blacklist, not both.", name)
			}

			validateThresholds(schema.DefaultThresholds, name)

			for _, table := range schema.TableOverrides {
				assertNotEmptyf(table.TableName, "table overrides for %s must specify a name", name)
				name := fmt.Sprintf("%s.%s", name, table.TableName)
				if table.LatencySpec != nil {
					validateThresholds(table.LatencySpec.Thresholds, name)
				}
			}
		}
	}
}

func validateThresholds(thresholds *models.Thresholds, name string) {
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
		assertNilf(err, "invalid threshold tier: %s in %s", tier, name)

		if threshold == helpers.NoLatencyAlert || threshold == "" {
			continue
		}

		_, err = time.ParseDuration(threshold)
		assertNilf(err, "invalid duration %s for tier: %s in %s", threshold, tier, name)
	}
}
