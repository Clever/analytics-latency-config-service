# analytics-latency-config-service

Service for surfacing latency config settings for the analytics pipeline.

Owned by eng-deip

This service is still a work in progress. It was split out from [Analytics Pipeline Monitor](https://github.com/Clever/analytics-pipeline-monitor/pulls] to consolidate the usage of the latency configs with other jobs.
However, APM still uses the raw configs to generate table checks.
In the future, we should consider doing the schema-level whitelist/blacklist (and threshold default) consolidation into ALCS, but that's out of scope right now.
We also need to be careful, since ALCS should not be taking a long time to determine true table-latency, and instead focus on the latency stats table for fast lookups.

## Developing

- Update swagger.yml with your endpoints. See the [Swagger spec](http://swagger.io/specification/) for additional details on defining your swagger file.

- Run `make generate` to generate the supporting code

- Run `make build`, `make run`, or `make test` - This should fail with an error about having to implement the business logic.

- Implement aforementioned business logic so that code will build

## Declaring New Latency Checks
Defining checks in analytics-latency-config-service can be accomplished by adding a new entry to `config/latency_config.json` in the following format:

```
  "prod": [
    {
      "schema": "mongo",
      "default_threshold": "24h",
      "default_timestamp_column": "_data_timestamp",
      "blacklist": ["billing_03_31_snapshot"],
      "table_overrides": [
        {
          "table": "districts",
          "latency": {
            "timestamp_column": "_data_timestamp",
            "thresholds": {
              "critical": "8h",
              "major": "4h",
              "minor": "2h"
            }
          }
        },
      ]
    }, ...
  ]
```

`analytics-latency-config-service` then reads from this config to surface latency info to other workers/services. `schema` + `table` identifies the table, and `latency.timestamp_column` identifies the time a row enters Redshift. `latency.thresholds` configures the different tiers of latency thresholds maximum amount of latency acceptable for the table's data in [Go time format](https://golang.org/pkg/time/#ParseDuration).

For tables that are not explicitly declared in the config, `default_threshold` and `default_timestamp_column` will be used as substitutes for the above values.

`whitelist` and `blacklist` allows tables to be whitelisted or blacklisted from latency checks. If neither is specified, all tables in the schema will be checked.

Note: `table_overrides`, `whitelist`, and `blacklist` must be in sync, as follows:

- If you include a table in `table_overrides`, it must also be listed in `whitelist`. (But not necessarily the other way around; you don't need to override everything you whitelist.)
- If you include a table in `blacklist`, it must NOT have an entry in `table_overrides`.

## Deploying

```
ark start analytics-latency-config-service
```
