package config

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// TestParseChecks verifies that the latency_config
// JSON can successfully be unmarshalled.
// If this test fails, it is likely that there is a
// formatting issue with the latency_config
func TestParseChecks(t *testing.T) {
	latencyConfigPath = "latency_config.json"
	assert.NotPanics(t, func() {
		ParseChecks()
	}, "Unable to parse latency checks")
}
