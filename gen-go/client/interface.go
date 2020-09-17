package client

import (
	"context"

	"github.com/Clever/analytics-latency-config-service/gen-go/models"
)

//go:generate mockgen -source=$GOFILE -destination=mock_client.go -package=client

// Client defines the methods available to clients of the analytics-latency-config-service service.
type Client interface {

	// HealthCheck makes a GET request to /_health
	// Checks if the service is healthy
	// 200: nil
	// 400: *models.BadRequest
	// 500: *models.InternalError
	// default: client side HTTP errors, for example: context.DeadlineExceeded.
	HealthCheck(ctx context.Context) error

	// GetAllLegacyConfigs makes a GET request to /legacy_config
	//
	// 200: *models.AnalyticsLatencyConfigs
	// 400: *models.BadRequest
	// 500: *models.InternalError
	// default: client side HTTP errors, for example: context.DeadlineExceeded.
	GetAllLegacyConfigs(ctx context.Context) (*models.AnalyticsLatencyConfigs, error)
}
