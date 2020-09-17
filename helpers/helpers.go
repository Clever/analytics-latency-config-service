package helpers

import (
	"fmt"

	"github.com/Clever/analytics-latency-config-service/gen-go/models"
)

const NoLatencyAlert = "none"

// CombineThresholdsWithDefaults takes some threshold overrides, and fills in default values if not specified
func CombineThresholdsWithDefaults(overrides *models.Thresholds, defaults models.Thresholds) models.Thresholds {
	if overrides == nil {
		return defaults
	}

	if overrides.Critical == "" {
		overrides.Critical = defaults.Critical
	}
	if overrides.Major == "" {
		overrides.Major = defaults.Major
	}
	if overrides.Minor == "" {
		overrides.Minor = defaults.Minor
	}

	return *overrides
}

// GetThresholdTierValue returns the threshold value for the latency type
func GetThresholdTierValue(threshold *models.Thresholds, tier models.ThresholdTier) (string, error) {
	switch tier {
	case models.ThresholdTierCritical:
		return threshold.Critical, nil
	case models.ThresholdTierMajor:
		return threshold.Major, nil
	case models.ThresholdTierMinor:
		return threshold.Minor, nil
	case models.ThresholdTierNone:
		fallthrough // there's no field for none on the thresholds object
	default:
		return NoLatencyAlert, fmt.Errorf("unexpected threshold tier")
	}
}

// GetThresholdTierErrorValue returns the error value to report for the latency error
func GetThresholdTierErrorValue(tier models.ThresholdTier) (int, error) {
	switch tier {
	case models.ThresholdTierCritical:
		return 3, nil
	case models.ThresholdTierMajor:
		return 2, nil
	case models.ThresholdTierMinor:
		return 1, nil
	case models.ThresholdTierNone:
		return 0, nil
	default:
		return 0, fmt.Errorf("unexpected threshold tier")
	}
}

func GetDatabaseConfig(configs models.AnalyticsLatencyConfigs, database models.AnalyticsDatabase) ([]*models.SchemaConfig, error) {
	switch database {
	case models.AnalyticsDatabaseRedshiftProd:
		return configs.RedshiftProd, nil
	case models.AnalyticsDatabaseRedshiftFast:
		return configs.RedshiftFast, nil
	case models.AnalyticsDatabaseRdsInternal:
		return configs.RdsInternal, nil
	case models.AnalyticsDatabaseRdsExternal:
		return configs.RdsExternal, nil
	default:
		return nil, fmt.Errorf("unexpected database config")
	}
}
