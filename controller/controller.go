package controller

import (
	"context"
	"fmt"
	"log"

	"github.com/Clever/analytics-latency-config-service/gen-go/models"
)

// Controller implements server.Controller
type Controller struct {
}

func New() (*Controller, error) {
	return &Controller{}, nil
}

// HealthCheck handles GET requests to /_health
func (c Controller) HealthCheck(ctx context.Context) error {
	return nil
}

// GetAllLegacyConfigs is a legacy function to support APM so that it can query all the configs
func (c Controller) GetAllLegacyConfigs(ctx context.Context) (*models.AnalyticsLatencyConfigs, error) {
	err := fmt.Errorf("Not yet implemented")
	log.Fatal(err)
	return nil, err
}
