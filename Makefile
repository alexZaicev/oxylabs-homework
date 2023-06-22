GO   := go
WIRE := wire

export GOFLAGS     := -mod=vendor

DIR_DIST := ./dist
DIR_CMD  := ./cmd

VERSION = $(shell (git describe --long --tags --match 'v[0-9]*' || echo v0.0.0) | cut -c2-)

#=================================================
# Service build targets
#=================================================

.PHONY: build
build:
	@mkdir -p $(DIR_DIST)
	CGO_ENABLED=0 $(GO) build -ldflags="-X main.version=$(VERSION)" -o $(DIR_DIST)/ $(DIR_CMD)/...

.PHONY: vendor
vendor:
	$(GO) mod vendor

.PHONY: wire
wire: wire-server wire-publisher wire-subscriber

.PHONY: wire-%
wire-%:
	$(eval TARGET := $(subst wire-,,$@))
	$(WIRE) $(DIR_CMD)/$(TARGET)