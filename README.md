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

## Local docker-compose

### Commands

- Start all services: `docker-compose up -d`
- Stop all services: `docker-compose stop`
- Destroy all services: `docker-compose down`

### Postgres Credentials

pgAdmin URL: `http://localhost:5050`

**pgAdmin 4 Credentials**

- Username: `user@domain.com`
- Password: `SuperSecret`

**Postgres Connection Credentials**

- Host: `postgres` (since it's docker-to-docker communications)
- Username: `postgres`
- Password: `example`

_Note: May need to manually destroy `postgres` docker container upon credentials change_
