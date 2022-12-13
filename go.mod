module github.com/Clever/analytics-latency-config-service

go 1.16

require (
	github.com/Clever/analytics-latency-config-service/gen-go/models v0.0.0-00010101000000-000000000000
	github.com/Clever/go-process-metrics v0.4.0
	github.com/Clever/kayvee-go/v7 v7.6.0
	github.com/Clever/pq v0.2.0
	github.com/Clever/wag v4.1.0+incompatible
	github.com/ghodss/yaml v1.0.0
	github.com/go-errors/errors v1.1.1
	github.com/go-openapi/strfmt v0.21.2
	github.com/go-openapi/swag v0.21.1
	github.com/golang/mock v1.6.0
	github.com/gorilla/handlers v1.5.1
	github.com/gorilla/mux v1.8.0
	github.com/hashicorp/go-multierror v1.1.0
	github.com/kardianos/osext v0.0.0-20190222173326-2bc1f35cddc0
	github.com/kevinburke/go-bindata v3.21.0+incompatible
	github.com/snowflakedb/gosnowflake v1.6.3
	github.com/stretchr/testify v1.8.0
	github.com/xeipuuv/gojsonschema v1.2.1-0.20200424115421-065759f9c3d7 // indirect
	go.opentelemetry.io/contrib/instrumentation/github.com/gorilla/mux/otelmux v0.34.0
	go.opentelemetry.io/otel v1.9.0
	go.opentelemetry.io/otel/exporters/otlp/otlptrace v1.9.0
	go.opentelemetry.io/otel/exporters/otlp/otlptrace/otlptracegrpc v1.9.0
	go.opentelemetry.io/otel/sdk v1.9.0
	go.opentelemetry.io/otel/trace v1.9.0
	golang.org/x/xerrors v0.0.0-20220609144429-65e65417b02f
	gopkg.in/mgo.v2 v2.0.0-20190816093944-a6b53ec6cb22 // indirect
)

exclude (
	github.com/codahale/hdrhistogram v1.0.0
	github.com/codahale/hdrhistogram v1.0.1
	github.com/codahale/hdrhistogram v1.1.0
)

// Do not delete: the following line allows for the server module to import the local version of the models without first publishing a new models module
replace github.com/Clever/analytics-latency-config-service/gen-go/models => ./gen-go/models
