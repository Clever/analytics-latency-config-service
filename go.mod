module github.com/Clever/analytics-latency-config-service

go 1.13

require (
	github.com/Clever/discovery-go v1.7.2
	github.com/Clever/go-process-metrics v0.2.0
	github.com/Clever/pq v0.0.0-20210406222402-741030d37ece
	github.com/Clever/wag v1.6.1
	github.com/afex/hystrix-go v0.0.0-20180502004556-fa1af6a1f4f5
	github.com/asaskevich/govalidator v0.0.0-20200817114649-df4adffc9d8c // indirect
	github.com/codahale/hdrhistogram v0.9.0 // indirect
	github.com/donovanhide/eventsource v0.0.0-20171031113327-3ed64d21fb0b
	github.com/ghodss/yaml v1.0.0
	github.com/go-errors/errors v1.1.1
	github.com/go-openapi/errors v0.19.6
	github.com/go-openapi/runtime v0.19.21 // indirect
	github.com/go-openapi/spec v0.19.9 // indirect
	github.com/go-openapi/strfmt v0.19.5
	github.com/go-openapi/swag v0.19.9
	github.com/go-openapi/validate v0.19.10
	github.com/golang/mock v1.6.0
	github.com/gorilla/handlers v1.5.1
	github.com/gorilla/mux v1.8.0
	github.com/hashicorp/go-multierror v1.1.0
	github.com/kardianos/osext v0.0.0-20190222173326-2bc1f35cddc0
	github.com/kevinburke/go-bindata v3.21.0+incompatible
	github.com/mailru/easyjson v0.7.6 // indirect
	github.com/mitchellh/mapstructure v1.3.3 // indirect
	github.com/opentracing/opentracing-go v1.2.0
	github.com/smartystreets/goconvey v1.6.4 // indirect
	github.com/snowflakedb/gosnowflake v1.6.3
	github.com/stretchr/testify v1.7.0
	github.com/uber/jaeger-client-go v2.25.0+incompatible
	github.com/uber/jaeger-lib v2.2.0+incompatible // indirect
	github.com/xeipuuv/gojsonpointer v0.0.0-20190905194746-02993c407bfb // indirect
	github.com/xeipuuv/gojsonschema v1.2.1-0.20200424115421-065759f9c3d7 // indirect
	go.mongodb.org/mongo-driver v1.4.1 // indirect
	go.uber.org/atomic v1.7.0 // indirect
	golang.org/x/net v0.0.0-20210813160813-60bc85c4be6d
	golang.org/x/xerrors v0.0.0-20200804184101-5ec99f83aff1
	gopkg.in/Clever/kayvee-go.v6 v6.24.0
	gopkg.in/mgo.v2 v2.0.0-20190816093944-a6b53ec6cb22 // indirect
	gopkg.in/yaml.v2 v2.3.1-0.20200602174213-b893565b90ca // indirect
)

exclude (
	github.com/codahale/hdrhistogram v1.0.0
	github.com/codahale/hdrhistogram v1.0.1
	github.com/codahale/hdrhistogram v1.1.0
)
