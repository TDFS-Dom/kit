GOPATH:=$(shell go env GOPATH)
VERSION=$(shell git describe --tags --always)
BUF_VERSION=v1.8.0
DIR=$(shell pwd)

SRV_PROTO_DIR = dtm event oidc user sys saas realtime
PKG_PROTO_DIR = $(patsubst %/,%,$(shell cd pkg && ls -d */))
OTHER_PROTO_DIR = $(patsubst %/,%,$(shell cd proto && ls -d */))

.PHONY: link
# link proto
link:
	for d in $(SRV_PROTO_DIR); do \
  		ln -sfn $(DIR)/$$d $(DIR)/buf/$$d; \
  	done
	for d in $(PKG_PROTO_DIR); do \
		ln -sfn $(DIR)/pkg/$$d $(DIR)/buf/$$d; \
    done
	for d in $(OTHER_PROTO_DIR); do \
		ln -sfn $(DIR)/proto/$$d $(DIR)/buf/$$d; \
    done

.PHONY: init
# init env
init:
	make link
	go install github.com/go-kratos/kratos/cmd/kratos/v2@latest
	go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
	go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest
	go install github.com/go-saas/kit/cmd/protoc-gen-go-grpc-proxy@ffede0b182a5c92e22da30585ba5a74658aff773
	go install github.com/go-kratos/kratos/cmd/protoc-gen-go-http/v2@latest
	go install github.com/go-kratos/kratos/cmd/protoc-gen-go-errors/v2@latest
	go install github.com/go-saas/kit/cmd/protoc-gen-go-errors-i18n/v2@ffede0b182a5c92e22da30585ba5a74658aff773
	go install github.com/envoyproxy/protoc-gen-validate@v0.6.7
	go install github.com/bufbuild/buf/cmd/buf@$(BUF_VERSION)
	go install github.com/bufbuild/buf/cmd/protoc-gen-buf-breaking@$(BUF_VERSION)
	go install github.com/bufbuild/buf/cmd/protoc-gen-buf-lint@$(BUF_VERSION)

.PHONY: user
user:
	make api
	cd user && $(MAKE) all

.PHONY: saas
saas:
	make api
	cd saas && $(MAKE) all

.PHONY: realtime
realtime:
	make api
	cd realtime && $(MAKE) all

.PHONY: sys
sys:
	make api
	cd sys && $(MAKE) all

.PHONY: apisix
apisix:
	cd gateway/apisix && $(MAKE) all


.PHONY: api
# generate api proto
api:
	buf generate --path ./buf/user --path ./buf/sys --path ./buf/saas --path ./buf/realtime --path ./buf/dtm --path ./buf/event --path ./buf/oidc
	cd user && $(MAKE) api
	cd saas && $(MAKE) api
	cd sys && $(MAKE) api
	cd realtime && $(MAKE) api

.PHONY: generate
# generate
generate:
	go generate ./pkg/...

.PHONY: examples
# generate
examples:
	cd examples/monolithic && $(MAKE) all

.PHONY: build
build:
	cd user && $(MAKE) build
	cd saas && $(MAKE) build
	cd sys && $(MAKE) build
	cd realtime && $(MAKE) build
	cd gateway/apisix && $(MAKE) build
	cd examples/monolithic && $(MAKE) build


.PHONY: all
all:
	go mod tidy
	make generate
	make api
	make user
	make saas
	make sys
	make realtime
	make apisix
	make examples

# show help
help:
	@echo ''
	@echo 'Usage:'
	@echo ' make [target]'
	@echo ''
	@echo 'Targets:'
	@awk '/^[a-zA-Z\-\_0-9]+:/ { \
	helpMessage = match(lastLine, /^# (.*)/); \
		if (helpMessage) { \
			helpCommand = substr($$1, 0, index($$1, ":")-1); \
			helpMessage = substr(lastLine, RSTART + 2, RLENGTH); \
			printf "\033[36m%-22s\033[0m %s\n", helpCommand,helpMessage; \
		} \
	} \
	{ lastLine = $$0 }' $(MAKEFILE_LIST)

.DEFAULT_GOAL := all