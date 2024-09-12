package helpers

import (
	"fmt"
	"time"

	"github.com/Clever/analytics-latency-config-service/gen-go/models"
)

const (
	// NoLatencyAlert is a constant to define the string used when there's no latency alert configured for a given threshold
	NoLatencyAlert = "none"
)

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
	if overrides.Refresh == "" {
		overrides.Refresh = defaults.Refresh
	}

	return *overrides
}

// GetThresholdTierValue returns the threshold value for the latency type
func GetThresholdTierValue(thresholds *models.Thresholds, tier models.ThresholdTier) (string, error) {
	if thresholds == nil {
		return NoLatencyAlert, fmt.Errorf("cannot get %s threshold value from a nil thresholds", tier)
	}
	switch tier {
	case models.ThresholdTierCritical:
		return thresholds.Critical, nil
	case models.ThresholdTierMajor:
		return thresholds.Major, nil
	case models.ThresholdTierMinor:
		return thresholds.Minor, nil
	case models.ThresholdTierRefresh:
		return thresholds.Refresh, nil
	case models.ThresholdTierNone:
		fallthrough // there's no field for none on the thresholds object
	default:
		return NoLatencyAlert, fmt.Errorf("unexpected threshold tier: %s", tier)
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
	case models.ThresholdTierRefresh:
		fallthrough // Refreshes tier isn't an error
	case models.ThresholdTierNone:
		return 0, nil
	default:
		return 0, fmt.Errorf("unexpected threshold tier: %s", tier)
	}
}

// GetDatabaseConfig returns the database-specific config
func GetDatabaseConfig(configs models.AnalyticsLatencyConfigs, database models.AnalyticsDatabase) ([]*models.SchemaConfig, error) {
	switch database {
	case models.AnalyticsDatabaseRedshiftFast:
		return configs.RedshiftFast, nil
	case models.AnalyticsDatabaseRdsExternal:
		return configs.RdsExternal, nil
	case models.AnalyticsDatabaseSnowflake:
		return configs.Snowflake, nil
	default:
		return nil, fmt.Errorf("unexpected database config: %s", database)
	}
}

// CheckThresholdCrossed checks if a latency crosses a particular tier for the latency thresholds
func CheckThresholdCrossed(latency float64, thresholds *models.Thresholds, tier models.ThresholdTier) (bool, string, error) {
	limit, err := GetThresholdTierValue(thresholds, tier)
	if err != nil {
		return false, "", err
	}

	// Override to not check threshold tier
	if limit == NoLatencyAlert {
		return false, NoLatencyAlert, nil
	}

	duration, err := time.ParseDuration(limit)
	if err != nil {
		return false, "", fmt.Errorf("could not parse duration %s: %s", limit, err.Error())
	}

	if latency > duration.Hours() {
		return true, limit, nil
	}
	return false, limit, nil
}
