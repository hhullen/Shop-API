MOCKGEN_INSTALL=go install github.com/golang/mock/mockgen@latest
MOCKGEN_BIN=$(shell where mockgen)

SQLC_INSTALL=go install github.com/sqlc-dev/sqlc/cmd/sqlc@latest
SQLC_BIN=$(shell where sqlc)

SWAG_INSTALL=go install github.com/swaggo/swag/cmd/swag@latest
SWAG_BIN=$(shell where swag)
SWAG_DOCS_DIR=internal/docs

SHADOW_INSTALL=golang.org/x/tools/go/analysis/passes/shadow/cmd/shadow@latest
SHADOW_BIN=$(shell where shadow)

ERRCHECK_INSTALL=go install github.com/kisielk/errcheck@latest
ERRCHECK_BIN=$(shell where errcheck)

COVERAGE_FILE=coverage.out

SERVICE_DATASTRUCT_DIR=internal/datastruct
API_DIR=internal/api/v1

PWD=$(pwd)
RM=rm -rf

EXTENSION=.out

MIGRATOR_DIR=./cmd/migrator
MIGRATOR_BIN=$(MIGRATOR_DIR)/migrator$(EXTENSION)

SHOPAPI_DIR=./cmd/shopapi
SHOPAPI_BIN=$(SHOPAPI_DIR)/shopapi$(EXTENSION)

LOCAL_DB_NAME=shop-api-local-database
LOCAL_DB_DATA_NAME=local_shopapi_postgres_data

ifeq ($(OS),Windows_NT)
	SHELL=powershell.exe
	EXTENSION=.exe
	PWD=$(shell powershell -Command "(Get-Location).Path")
	RM=echo
	RM_POSTFIX=| Remove-Item -Force -ErrorAction SilentlyContinue; exit 0
	SQLC_BIN=$(strip $(shell (Get-Command sqlc.exe -ErrorAction SilentlyContinue).Source))
	MOCKGEN_BIN=$(strip $(shell (Get-Command mockgen.exe -ErrorAction SilentlyContinue).Source))
	SWAG_BIN=$(strip $(shell (Get-Command swag.exe -ErrorAction SilentlyContinue).Source))
	SHADOW_BIN=$(strip $(shell (Get-Command shadow.exe -ErrorAction SilentlyContinue).Source))
	ERRCHECK_BIN=$(strip $(shell (Get-Command errcheck.exe -ErrorAction SilentlyContinue).Source))
endif

.PHONY: deps generate-sqlc generate-mocks generage-swag migrations-up migrations-down migrations-status start-local-database stop-local-database clean-local-database service coverage-info coverage-html

deps:
	go mod download

errcheck:
ifeq ($(ERRCHECK_BIN),)
	$(ERRCHECK_INSTALL)
endif
	errcheck.exe -verbose -ignoregenerated ./...

linter:
ifeq ($(SHADOW_BIN),)
	$(SHADOW_INSTALL)
endif
	go vet -vettool=$(SHADOW_BIN) ./...

generate-sqlc:
ifeq ($(SQLC_BIN),)
	$(SQLC_INSTALL)
endif
	sqlc generate

generate-mocks:
ifeq ($(MOCKGEN_BIN),)
	$(MOCKGEN_INSTALL)
endif
	go generate ./...

generate-swag:
ifeq ($(SWAG_BIN),)
	$(SWAG_INSTALL)
endif
	swag init -d $(SHOPAPI_DIR),$(SERVICE_DATASTRUCT_DIR),$(API_DIR) -o $(SWAG_DOCS_DIR)

$(MIGRATOR_BIN):
	go build -o $(MIGRATOR_BIN) $(MIGRATOR_DIR)

migrations-up: $(MIGRATOR_BIN)
	$(MIGRATOR_BIN) up

migrations-down: $(MIGRATOR_BIN)
	$(MIGRATOR_BIN) down

migrations-status: $(MIGRATOR_BIN)
	$(MIGRATOR_BIN) status

start-local-database:
	docker run -d --rm -p 5432:5432 -e POSTGRES_PASSWORD_FILE=/run/secrets/db_password -e POSTGRES_USER_FILE=/run/secrets/db_user -e POSTGRES_DB_FILE=/run/secrets/db_name -v $(PWD)/secrets/db_password.txt:/run/secrets/db_password:ro -v $(PWD)/secrets/db_user.txt:/run/secrets/db_user:ro -v $(PWD)/secrets/db_name.txt:/run/secrets/db_name:ro -v local_shopapi_postgres_data:/var/lib/postgresql/data --name $(LOCAL_DB_NAME) postgres:17.5-alpine3.21

stop-local-database:
	docker container stop $(LOCAL_DB_NAME)

clean-local-database:
	docker volume rm $(LOCAL_DB_DATA_NAME)

run-local: $(SHOPAPI_BIN)
	$(SHOPAPI_BIN)

run-local-fast:
	go run $(SHOPAPI_DIR)/main.go

$(SHOPAPI_BIN):
	go build -o $(SHOPAPI_BIN) $(SHOPAPI_DIR)

service:
	docker compose up

stop-service:
	docker compose down

service-rebuild:
	docker compose down -v
	docker compose up --build --renew-anon-volumes --force-recreate

$(COVERAGE_FILE):
	go test "-coverpkg=./..." "-coverprofile=coverage.out" ./...

coverage-info: $(COVERAGE_FILE)
	go tool cover "-func=coverage.out"

coverage-html: $(COVERAGE_FILE)
	go tool cover "-html=coverage.out"

clean:
	$(RM) $(MIGRATOR_BIN) $(SHOPAPI_BIN) $(COVERAGE_FILE) $(RM_POSTFIX)
