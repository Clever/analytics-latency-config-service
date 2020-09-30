package server

import (
	"context"

	"github.com/Clever/analytics-latency-config-service/gen-go/models"
)

//go:generate mockgen -source=$GOFILE -destination=mock_controller.go -package=server

// Controller defines the interface for the analytics-latency-config-service service.
type Controller interface {

	// HealthCheck handles GET requests to /_health
	// Checks if the service is healthy
	// 200: nil
	// 400: *models.BadRequest
	// 500: *models.InternalError
	// default: client side HTTP errors, for example: context.DeadlineExceeded.
	HealthCheck(ctx context.Context) error

	// GetTableLatency handles GET requests to /latency
	//
	// 200: *models.GetTableLatencyResponse
	// 400: *models.BadRequest
	// 404: *models.NotFound
	// 500: *models.InternalError
	// default: client side HTTP errors, for example: context.DeadlineExceeded.
	GetTableLatency(ctx context.Context, i *models.GetTableLatencyRequest) (*models.GetTableLatencyResponse, error)

	// GetAllLegacyConfigs handles GET requests to /legacy_config
	//
	// 200: *models.AnalyticsLatencyConfigs
	// 400: *models.BadRequest
	// 500: *models.InternalError
	// default: client side HTTP errors, for example: context.DeadlineExceeded.
	GetAllLegacyConfigs(ctx context.Context) (*models.AnalyticsLatencyConfigs, error)
}
