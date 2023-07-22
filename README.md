# intro-sqlc

## Environment

- Go 1.20.6
- sqlc v1.19.1

## Run

- Run PostgreSQL

```sh
cd db
docker compose up -d
```

- Run application

```sh
cd app
go run main.go
```

## Note

### Init Go Module

```sh
mkdir intro-sqlc
cd intro-sqlc
go mod init github.com/nukopy/intro-sqlc
```

### How to Use sqlc

1. Write SQL queries
2. Write `sqlc.yaml`
3. Execute `sqlc generate`
4. Write application code

### Communicate with PostgreSQL Container

- Enter PostgreSQL container

```sh
docker compose exec db bash
```

- Connect PostgreSQL REPL

```sh
docker compose exec db psql -d sample_db -U myuser
```

- Basic commands in PostgreSQL REPL

```sh
\du                   : show users
select current_user;  : show current user
\l                    : show databases
\c                    : show current database
\d                    : show relations (tables) on specified database
```
