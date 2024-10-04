COMMON_PACKAGE=github.com/zuu-development/fullstack-examination-2024/internal/common
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
  -X ${COMMON_PACKAGE}.version=${VERSION} \
  -X ${COMMON_PACKAGE}.buildDate=${BUILD_DATE} \
  -X ${COMMON_PACKAGE}.gitCommit=${GIT_COMMIT} \

.PHONY: cli
cli:
	GOOS=${HOST_OS} GOARCH=${HOST_ARCH} make cli-local

.PHONY: cli-local
cli-local:
	GODEBUG="tarinsecurepath=0,zipinsecurepath=0" go build -gcflags="all=-N -l" $(COVERAGE_FLAG) -v -ldflags '${LDFLAGS}' -o ${DIST_DIR}/${CLI_NAME} ./

.PHONY: ui
ui:
	cd ui && yarn build

.PHONY: dep-backend-local
dep-backend-local:
	go mod download

.PHONY: migrate
migrate:
	go run main.go migrate --config config.yaml

.PHONY: reset-local-db
reset-local-db:
	rm ./tmp/gorm.db

.PHONY: serve-backend
serve-backend:
	air -c .air.toml

.PHONY: dep-ui-local
dep-ui-local:
	cd ui && yarn install

.PHONY: serve-ui
serve-ui:
	cd ui && yarn dev

.PHONY: lint
lint:
	golangci-lint run
	cd ui && yarn lint

.PHONY: fmt
fmt:
	go mod tidy
	golangci-lint run --fix
	swag fmt
	cd ui && yarn lint-fix

.PHONY: test-backend
test-backend:
	gotestsum --format=testname --rerun-fails

test-backend-ci:
	gotestsum --format=testname -- -cover -coverprofile=coverage.out ./...

# ビルド時にチェックする .go ファイル
SWAG_GO_FILES:=$(shell find internal/handler -type f -name '*.go' -print)

docs/swagger.yaml: main.go $(SWAG_GO_FILES)
	swag init

docs/swagger.html: docs/swagger.yaml
	npx @redocly/cli@1.25.3 build-docs -o docs/swagger.html docs/swagger.yaml

.PHONY: swagger
swagger: docs/swagger.html