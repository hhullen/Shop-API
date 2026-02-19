# Simple Shop API

Here is a REST API service written with Golang net/http

## Requirements
GNU Make 3.81, go1.24.2, Docker 28.1.1 (Docker Desktop)

## How to run local
1. Prepare local database
```shell
make start-local-database
make migrations-up
```

2. Run service
```shell
make run-local
```

## How to run Docker Compose
1. Run servce
```shell
make service
```


## Explore via OpenAPI
When service is running the OpenAPI page with all implemented endpoints is available at [http://localhost:8080/swagger/index.html](http://localhost:8080/swagger/index.html)

## Makefile targets
- `make deps` download all required libraries
- `make errcheck` scan code for some errors
- `make linter` run linter
- `make generate-sqlc` generates Go code by sql queries. sqlc config is in `sqlc.yaml` file
- `make generate-mocks` generate mocks for interfaces specified in `//go:generate ...` comments in any Go file
- `make generate-swag` generates OpenAPI docs from comments of endpoint functions
- `make migrations-up` applies all migrations
- `make migrations-down` rollbacks of last migration
- `make migrations-status` shows migrations status
- `make start-local-database` runs local database
- `make stop-local-database` stops local database
- `make clean-local-database` cleans local database
- `make run-local` compiles and runs service binary
- `make service-run` runs service via `go run` command
- `make service` runs docker compose service
- `make stop-service` stops docker compose service
- `make service-rebuild` rebuild docker compose image
- `make clean` removes compiled binaries
