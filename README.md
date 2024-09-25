Invia exercise.
====

Develop a scalable RESTful API that handles user management,
including user creation, retrieval, update, and deletion.

Application can be run as service.

## Run app as CLI application

```shell
export INVIA_CONFIG_PATH=config/local.yaml
```

```shell
go run  ./cmd/users/main.go
```

## Run services as docker containers

```shell
docker compose up -d
```

## Stop server with docker containers

```shell
docker-compose down
```

## Documentation of API (swagger)

See swagger documentation file [here](swagger/swagger.yaml)
