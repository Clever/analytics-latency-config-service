include golang.mk
include wag.mk

SHELL := /bin/bash
export PATH := $(PWD)/bin:$(PATH)
APP_NAME ?= analytics-latency-config-service
EXECUTABLE = $(APP_NAME)
PKG = github.com/Clever/$(APP_NAME)
PKGS := $(shell go list ./... | grep -v /vendor | grep -v /gen-go | grep -v /tools | grep -v /tools)

# Temporarily pin to wag 6.4.5 until after migrated to go mod and Go 1.16
WAG_VERSION := latest

$(eval $(call golang-version-check,1.24))

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
	SNOWFLAKE_USER=$(SNOWFLAKE_USER_OVERRIDE) \
	SNOWFLAKE_DATABASE="DEV_QUICKBLOCKS_DB" \
	SNOWFLAKE_WAREHOUSE="PROD_ENG_XS_WH" \
	SNOWFLAKE_ROLE="CLEVER_ENG_ROLE" \
	SNOWFLAKE_AUTHENTICATOR="externalbrowser" \
	bin/$(EXECUTABLE)

generate: wag-generate-deps
	$(call wag-generate-mod,./swagger.yml)
	go mod vendor

install_deps:
	go mod vendor
	go build -o bin/mockgen -mod=vendor github.com/golang/mock/mockgen
	go build -o bin/go-bindata -mod=vendor github.com/kevinburke/go-bindata/go-bindata
