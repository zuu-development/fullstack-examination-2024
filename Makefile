CMD_PACKAGE=github.com/zuu-development/fullstack-examination-2024/cmd
CURRENT_DIR=$(shell pwd)
DIST_DIR=${CURRENT_DIR}/dist
CLI_NAME=todo-cli

HOST_OS:=$(shell go env GOOS)
HOST_ARCH:=$(shell go env GOARCH)

VERSION=$(shell cat ${CURRENT_DIR}/VERSION)
BUILD_DATE:=$(if $(BUILD_DATE),$(BUILD_DATE),$(shell date -u +'%Y-%m-%dT%H:%M:%SZ'))
GIT_COMMIT:=$(if $(GIT_COMMIT),$(GIT_COMMIT),$(shell git rev-parse HEAD))

ifeq (${COVERAGE_ENABLED}, true)
# We use this in the cli-local target to enable code coverage for e2e tests.
COVERAGE_FLAG=-cover
else
COVERAGE_FLAG=
endif

override LDFLAGS += \
  -X ${CMD_PACKAGE}.version=${VERSION} \
  -X ${CMD_PACKAGE}.buildDate=${BUILD_DATE} \
  -X ${CMD_PACKAGE}.gitCommit=${GIT_COMMIT} \

.PHONY: cli
cli:
	GOOS=${HOST_OS} GOARCH=${HOST_ARCH} make cli-local

.PHONY: cli-local
cli-local:
	GODEBUG="tarinsecurepath=0,zipinsecurepath=0" go build -gcflags="all=-N -l" $(COVERAGE_FLAG) -v -ldflags '${LDFLAGS}' -o ${DIST_DIR}/${CLI_NAME} ./

.PHONY: serve-api
serve-api:
	air -c .air.toml

.PHONY: serve-ui
serve-ui:
	cd ui && yarn dev

.PHONY: lint
lint:
	golangci-lint run

.PHONY: lint-fix
lint-fix:
	golangci-lint run --fix

dep-ui-local:
	cd ui && yarn install
