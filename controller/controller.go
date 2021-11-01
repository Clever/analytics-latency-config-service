package controller

import (
	"context"
	"fmt"

	"github.com/Clever/analytics-latency-config-service/config"
	"github.com/Clever/analytics-latency-config-service/db"
	"github.com/Clever/analytics-latency-config-service/gen-go/models"
	"github.com/Clever/analytics-latency-config-service/helpers"

	"github.com/hashicorp/go-multierror"
)

// Controller implements server.Controller
type Controller struct {
	redshiftProdConnection db.DBClient
	redshiftFastConnection db.DBClient
	rdsExternalConnection  db.DBClient
	rdsInternalConnection  db.DBClient
	snowflakeConnection    db.DBClient
	configChecks           models.AnalyticsLatencyConfigs
}

func (c *Controller) getDatabaseConnection(database models.AnalyticsDatabase) (db.DBClient, error) {
	switch database {
	case models.AnalyticsDatabaseRedshiftProd:
		return c.redshiftProdConnection, nil
	case models.AnalyticsDatabaseRedshiftFast:
		return c.redshiftFastConnection, nil
	case models.AnalyticsDatabaseRdsInternal:
		return c.rdsInternalConnection, nil
	case models.AnalyticsDatabaseRdsExternal:
		return c.rdsExternalConnection, nil
	case models.AnalyticsDatabaseSnowflake:
		return c.snowflakeConnection, nil
	default:
		return nil, fmt.Errorf("unexpected database")
	}
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

	rdsExternalConnection, err := db.NewRDSExternalClient()
	if err != nil {
		mErrors = multierror.Append(mErrors, fmt.Errorf("rds-external-failed-init: %s", err.Error()))
	}

	snowflakeConnection, err := db.NewSnowflakeProdClient()
	if err != nil {
		mErrors = multierror.Append(mErrors, fmt.Errorf("snowflake-failed-init: %s", err.Error()))
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
		snowflakeConnection:    snowflakeConnection,
		configChecks:           configChecks,
	}, nil
}

// HealthCheck handles GET requests to /_health
func (c *Controller) HealthCheck(ctx context.Context) error {
	return nil
}

// GetAllLegacyConfigs is a legacy function to support APM so that it can query all the configs
func (c *Controller) GetAllLegacyConfigs(ctx context.Context) (*models.AnalyticsLatencyConfigs, error) {
	return &c.configChecks, nil
}

// CheckTableThreshold returns the thresholds for a particular table.
func (c *Controller) GetTableLatency(ctx context.Context, i *models.GetTableLatencyRequest) (*models.GetTableLatencyResponse, error) {
	schemas, err := helpers.GetDatabaseConfig(c.configChecks, i.Database)
	if err != nil {
		return nil, models.BadRequest{Message: fmt.Sprintf("invalid database to query %s", i.Database)}
	}
	dbConn, err := c.getDatabaseConnection(i.Database)
	if err != nil {
		return nil, models.BadRequest{Message: fmt.Sprintf("invalid database to query %s", i.Database)}
	}

	if i.Schema == nil || *i.Schema == "" {
		return nil, models.BadRequest{Message: fmt.Sprintf("missing schema")}
	}
	if i.Table == nil || *i.Table == "" {
		return nil, models.BadRequest{Message: fmt.Sprintf("missing table")}
	}

	var schema *models.SchemaConfig
	for _, s := range schemas {
		if s.SchemaName == *i.Schema {
			schema = s
		}
	}
	if schema == nil {
		return nil, models.NotFound{Message: fmt.Sprintf("database %s does not have a config for schema %s", i.Database, *i.Schema)}
	}

	// Confirm table isn't in the blacklist
	for _, t := range schema.Blacklist {
		if t == *i.Table {
			return nil, models.NotFound{Message: fmt.Sprintf("database/schema %s/%s has table %s blacklisted", i.Database, *i.Schema, *i.Table)}
		}
	}

	tableThresholds := helpers.CombineThresholdsWithDefaults(schema.DefaultThresholds, config.GlobalDefaultThresholds)
	tableOwner := config.DefaultOwner
	if schema.SchemaOwner != "" {
		tableOwner = schema.SchemaOwner
	}

	for _, t := range schema.TableOverrides {
		if t.TableName == *i.Table {
			if t.LatencySpec != nil {
				tableThresholds = helpers.CombineThresholdsWithDefaults(t.LatencySpec.Thresholds, tableThresholds)
			}
			if t.TableOwner != "" {
				tableOwner = t.TableOwner
			}
			break
		}
	}

	latency, found, err := dbConn.QueryLatencyTable(*i.Schema, *i.Table)
	if err != nil {
		return nil, models.InternalError{Message: fmt.Sprintf("unexpected error trying to query latency table for %s.%s.%s: %s", i.Database, *i.Schema, *i.Table, err.Error())}
	}

	reportLatency := (*float64)(nil)
	if found {
		// Silly Golang. Casting is for kids!
		floatLatency := (float64)(latency)
		reportLatency = &floatLatency
	}

	return &models.GetTableLatencyResponse{
		Database:   i.Database,
		Schema:     i.Schema,
		Table:      i.Table,
		Owner:      &tableOwner,
		Thresholds: &tableThresholds,
		Latency:    reportLatency,
	}, nil
}
