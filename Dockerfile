FROM alpine:3.10

RUN apk add ca-certificates
RUN update-ca-certificates

COPY kvconfig.yml /bin/kvconfig.yml
COPY bin/analytics-latency-config-service /bin/analytics-latency-config-service

CMD ["/bin/analytics-latency-config-service", "--addr=0.0.0.0:80"]
