include golang.mk
include wag.mk

SHELL := /bin/bash
export PATH := $(PWD)/bin:$(PATH)
APP_NAME ?= analytics-latency-config-service
EXECUTABLE = $(APP_NAME)
PKG = github.com/Clever/$(APP_NAME)
PKGS := $(shell go list ./... | grep -v /vendor | grep -v /gen-go | grep -v /tools)

WAG_VERSION := latest

$(eval $(call golang-version-check,1.13))

export POSTGRES_USER?=postgres
export POSTGRES_HOST?=localhost
export POSTGRES_PORT?=5432
export POSTGRES_DB?=alcs-integration

.PHONY: all test build run $(PKGS) generate install_deps

all: test build

test: $(PKGS)
$(PKGS): golang-test-all-deps
	$(call golang-test-all,$@)

db-setup:
	createdb -h localhost -U postgres alcs-integration
	psql -h localhost -p 5432 -U postgres alcs-integration -c "CREATE FUNCTION GETDATE() RETURNS timestamp with time zone AS 'SELECT now()' LANGUAGE sql;"

build: generate
	$(call golang-build,$(PKG),$(EXECUTABLE))

run: build
	bin/$(EXECUTABLE)

generate: wag-generate-deps
	$(call wag-generate,./swagger.yml,$(PKG))
	go generate ./...

install_deps: golang-dep-vendor-deps
	$(call golang-dep-vendor)
	go build -o bin/mockgen    ./vendor/github.com/golang/mock/mockgen
	go build -o bin/go-bindata ./vendor/github.com/jteeuwen/go-bindata/go-bindata
