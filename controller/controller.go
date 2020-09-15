package controller

import (
	"context"
	"fmt"

	"github.com/Clever/analytics-latency-config-service/config"
	"github.com/Clever/analytics-latency-config-service/db"
	"github.com/Clever/analytics-latency-config-service/gen-go/models"

	"github.com/hashicorp/go-multierror"
)

// Controller implements server.Controller
type Controller struct {
	redshiftProdConnection db.PostgresClient
	redshiftFastConnection db.PostgresClient
	rdsExternalConnection  db.PostgresClient
	rdsInternalConnection  db.PostgresClient
	configChecks           models.AnalyticsLatencyConfigs
}

func New() (*Controller, error) {
	var mErrors *multierror.Error
	redshiftProdConnection, err := db.NewRedshiftProdClient()
	if err != nil {
		mErrors = multierror.Append(mErrors, fmt.Errorf("redshift-prod-failed-init: %s", err.Error()))
	}

	redshiftFastConnection, err := db.NewRedshiftFastClient()
	if err != nil {
		mErrors = multierror.Append(mErrors, fmt.Errorf("redshift-fast-failed-init: %s", err.Error()))
	}

	rdsInternalConnection, err := db.NewRDSInternalClient()
	if err != nil {
		mErrors = multierror.Append(mErrors, fmt.Errorf("rds-internal-failed-init: %s", err.Error()))
	}

	rdsExternalConnection, err := db.NewRDSInternalClient()
	if err != nil {
		mErrors = multierror.Append(mErrors, fmt.Errorf("rds-external-failed-init: %s", err.Error()))
	}

	configChecks := config.ParseChecks()

	if v := mErrors.ErrorOrNil(); v != nil {
		return nil, v
	}

	return &Controller{
		redshiftProdConnection: redshiftProdConnection,
		redshiftFastConnection: redshiftFastConnection,
		rdsExternalConnection:  rdsExternalConnection,
		rdsInternalConnection:  rdsInternalConnection,
		configChecks:           configChecks,
	}, nil
}

// HealthCheck handles GET requests to /_health
func (c Controller) HealthCheck(ctx context.Context) error {
	return nil
}

// GetAllLegacyConfigs is a legacy function to support APM so that it can query all the configs
func (c Controller) GetAllLegacyConfigs(ctx context.Context) (*models.AnalyticsLatencyConfigs, error) {
	return &c.configChecks, nil
}
