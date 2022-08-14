.PHONY: build run run-api doc validate spec clean web help

all: run-sqlc validate clean build

validate:
	swagger validate ./api/v1/swagger.yml

spec:
	swagger generate spec -o ./api/v1/swagger-gen.yml

build:
	swagger -q generate server -A sms-microservice -f api/v1/swagger.yml -s internal/apis -m internal/models
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags="-w -s" -v ./cmd/sms-microservice-server


api: validate clean build
	./sms-microservice-server --port=8081 --host=0.0.0.0

run: all
	./sms-microservice-server --port=8081 --host=0.0.0.0

doc:
	swagger validate ./api/v1/swagger.yml
	swagger serve api/v1/swagger.yml --no-open --host=0.0.0.0 --port=9090 --base-path=/

clean:
	rm -rf sms-microservice-server
	go clean -i .

run-sqlc:
	./run-sqlc.sh

migrate-up:
	migrate -path internal/repository/schema -database "mysql://maul:maul@tcp(52.77.253.103:3306)/sms-microservice?query" -verbose up

migrate-down:
	migrate -path internal/repository/schema -database "mysql://maul:maul@tcp(127.0.0.1:3306)/sms-microservice?query" -verbose down

create-migration:
	migrate create -ext sql -dir internal/repository/schema -seq $(action)

help:
	@echo "make: compile packages and dependencies"
	@echo "make validate: OpenAPI validation"
	@echo "make spec: OpenAPI Spec"
	@echo "make clean: remove object files and cached files"
	@echo "make build: Generate Server and Client API"
	@echo "make doc: Serve the Doc UI"
	@echo "make web: Build svelte static"
	@echo "make run-api: Build api only and serve"
	@echo "make run: Build everything and serve"
