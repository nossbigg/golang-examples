# golang-examples

## Formatting code

```bash
gofmt -w .
```

## Running locally

- Clone repo within `$GOPATH`

```bash
git clone git@github.com:nossbigg/golang-examples.git "$GOPATH"/src/nossbigg.com/golangexamples
```

## Running tests

```bash
go test ./...
```

## docker-compose Commands

- Start all services: `docker-compose up -d`
- Stop all services: `docker-compose stop`
- Destroy all services: `docker-compose down`

## Postgres Credentials

pgAdmin URL: `http://localhost:5050`

**pgAdmin 4 Credentials**

- Username: `user@domain.com`
- Password: `SuperSecret`

**Postgres Connection Credentials**

- Host: `postgres` (since it's docker-to-docker communications)
- Username: `postgres`
- Password: `example`

_Note: May need to manually destroy `postgres` docker container upon credentials change_

## Redis Credentials

**redisinsight Setup**

- Select Add Redis Database
- Host: `redis`
- Port: `6379`
- Name: `nossbigg`

## Prerequisites & Related Documentation

Languages

- [Golang](https://golang.org/)

Containerization

- [Docker](https://www.docker.com/products/docker-desktop)

DB & Cache

- [PostgreSQL](https://www.postgresql.org/)
- [pgAdmin](https://www.pgadmin.org/)
- [Redis](https://redis.io/)
- [redisinsight](https://redislabs.com/redis-enterprise/redis-insight/)

CLI tools

- [golang-migrate](https://github.com/golang-migrate/migrate)
- [redis-cli](https://redis.io/topics/rediscli)
- [make](https://www.gnu.org/software/make/)
- [PostgreSQL CLI tools](https://formulae.brew.sh/formula/postgresql)
