dep:
	go mod tidy

gen:
	oapi-codegen --old-config-style --generate types -o api/restapi/openapi_types.gen.go --package restapi ./swagger/swagger.yaml
	oapi-codegen --old-config-style --generate gorilla -o api/restapi/openapi_server.gen.go --package restapi ./swagger/swagger.yaml

build:
	go build -o users ./cmd/users/main.go

buildstatic:
	go build -tags musl -ldflags="-w -extldflags '-static' " -o users ./cmd/users/main.go

run:
	go run  ./cmd/users/main.go

test:
	go test ./...

lint:
	golangci-lint run ./...

check-migrate:
	which migrate || (go install -tags 'postgres' github.com/golang-migrate/migrate/v4/cmd/migrate@latest)

migrate: check-migrate
	migrate -database postgres://${POSTGRES_USER}:${POSTGRES_PASSWORD}@${POSTGRES_HOST}:${POSTGRES_PORT}/${POSTGRES_DB}?sslmode=disable -path migrations up

migrate-down: check-migrate
	migrate -database postgres://${POSTGRES_USER}:${POSTGRES_PASSWORD}@${POSTGRES_HOST}:${POSTGRES_PORT}/${POSTGRES_DB}?sslmode=disable -path migrations down $(V)	