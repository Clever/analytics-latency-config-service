# analytics-latency-config-service

Service for surfacing latency config settings for the analytics pipeline.

Owned by eng-deip

## Developing

- Update swagger.yml with your endpoints. See the [Swagger spec](http://swagger.io/specification/) for additional details on defining your swagger file.

- Run `make generate` to generate the supporting code

- Run `make build`, `make run`, or `make test` - This should fail with an error about having to implement the business logic.

- Implement aforementioned business logic so that code will build

## Deploying

```
ark start analytics-latency-config-service
```
