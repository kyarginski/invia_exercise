Invia exercise.
====

Develop a scalable RESTful API that handles user management,
including user creation, retrieval, update, and deletion.

Application can be run as service.

## Build application

Build and check code can be done with the `make` command via Makefile.


## Run application as CLI application

```shell
export INVIA_CONFIG_PATH=config/local.yaml
```

```shell
go run  ./cmd/users/main.go
```

## Run application as docker containers

```shell
docker compose up -d
```

## Stop application with docker containers

```shell
docker compose down
```

## Documentation of API (swagger)

See swagger documentation file [here](swagger/swagger.yaml)

Swagger file can be used for API testing.

## Using OpenTelemetry

https://www.jaegertracing.io/docs/1.47/getting-started/

Start (see docker-compose.yml)
```
docker run -d --name jaeger \
  -e COLLECTOR_ZIPKIN_HOST_PORT=:9411 \
  -e COLLECTOR_OTLP_ENABLED=true \
  -p 6831:6831/udp \
  -p 6832:6832/udp \
  -p 5778:5778 \
  -p 16686:16686 \
  -p 4317:4317 \
  -p 4318:4318 \
  -p 14250:14250 \
  -p 14268:14268 \
  -p 14269:14269 \
  -p 9411:9411 \
  jaegertracing/all-in-one:1.61.0

```

See results

http://localhost:16686 to access the Jaeger UI.


Add request-id into request header

```
request-id=fe1f3f07-8eb3-11ee-829b-0242ac130006
```

## Run tests

```shell
make test
```

## Test coverage

![Coverage](https://img.shields.io/badge/coverage-30.8%25-brightgreen)

To update test coverage, run:

```shell
make update-readme
```

## Implementation

1) The implementation starts with a contract description in the OpenAPI-format
`swagger/swagger.yaml`. 
2) According to the contract description, the code for the server is generated:
```shell
make gen
```
3) Next, we need to implement the implementation code in the file `api/impl.go` and the corresponding tests.
4) Generation of mocks for the test is performed by the command
```shell
make mocks
```
5) Testing can be performed using a contract file `swagger/swagger.yaml` or a utility program like `Postman`.
6) OpenTelemetry (Jaeger) can be used to check performance and errors in the application.

## Authentication and Authorization

For further improvements.

An application can use JWT for authentication and authorization. 
The JWT token will be generated when the user logs in and will be used for subsequent requests. 
The token will be stored in the `Authorization` header.

Or we can use a third-party system such as [KeyCloak](https://www.keycloak.org/)
