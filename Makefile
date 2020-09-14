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

.PHONY: all test build run $(PKGS) generate install_deps

all: test build

test: $(PKGS)
$(PKGS): golang-test-all-strict-deps
	$(call golang-test-all-strict,$@)

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
