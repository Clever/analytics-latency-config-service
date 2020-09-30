package helpers

import (
	"testing"

	"github.com/Clever/analytics-latency-config-service/gen-go/models"

	"github.com/stretchr/testify/assert"
)

func TestDefaultTableOverrides(t *testing.T) {
	assertions := assert.New(t)

	tests := []struct {
		title string

		// Input
		defaults       models.Thresholds
		TableOverrides *models.Thresholds

		// Output/behavior
		expectedThresholds models.Thresholds
	}{
		{
			title: "All thresholds overridden TableOverrides all",
			defaults: models.Thresholds{
				Critical: "9h",
				Major:    "5h",
				Minor:    "1h",
			},
			TableOverrides: &models.Thresholds{
				Critical: "8h",
				Major:    "4h",
				Minor:    "2h",
			},
			expectedThresholds: models.Thresholds{
				Critical: "8h",
				Major:    "4h",
				Minor:    "2h",
			},
		},
		{
			title: "No thresholds uses defaults",
			defaults: models.Thresholds{
				Critical: "9h",
				Major:    "5h",
				Minor:    "1h",
			},
			TableOverrides: &models.Thresholds{},
			expectedThresholds: models.Thresholds{
				Critical: "9h",
				Major:    "5h",
				Minor:    "1h",
			},
		},
		{
			title: "Nil thresholds uses defaults",
			defaults: models.Thresholds{
				Critical: "9h",
				Major:    "5h",
				Minor:    "1h",
			},
			TableOverrides: nil,
			expectedThresholds: models.Thresholds{
				Critical: "9h",
				Major:    "5h",
				Minor:    "1h",
			},
		},
		{
			title: "Empty thresholds are not overridden",
			defaults: models.Thresholds{
				Critical: "9h",
				Major:    "5h",
				Minor:    "1h",
			},
			TableOverrides: &models.Thresholds{
				Critical: "8h",
				Minor:    "2h",
			},
			expectedThresholds: models.Thresholds{
				Critical: "8h",
				Major:    "5h",
				Minor:    "2h",
			},
		},
		{
			title: "'none' value is not overridden",
			defaults: models.Thresholds{
				Critical: "9h",
				Major:    "5h",
				Minor:    "1h",
			},
			TableOverrides: &models.Thresholds{
				Critical: "",
				Major:    NoLatencyAlert,
				Minor:    "",
			},
			expectedThresholds: models.Thresholds{
				Critical: "9h",
				Major:    NoLatencyAlert,
				Minor:    "1h",
			},
		},
		{
			title: "default 'none' values are still overridden",
			defaults: models.Thresholds{
				Critical: NoLatencyAlert,
				Major:    NoLatencyAlert,
				Minor:    NoLatencyAlert,
			},
			TableOverrides: &models.Thresholds{
				Critical: "8h",
				Major:    "4h",
				Minor:    "",
			},
			expectedThresholds: models.Thresholds{
				Critical: "8h",
				Major:    "4h",
				Minor:    NoLatencyAlert,
			},
		},
	} // end tests definition

	for _, test := range tests {
		t.Logf("Testing DefaultTableOverrides %s", test.title)
		result := CombineThresholdsWithDefaults(test.TableOverrides, test.defaults)

		assertions.True(result == test.expectedThresholds, "TableOverrides match as expected")
	}

}
