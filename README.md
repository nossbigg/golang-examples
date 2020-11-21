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

### Credentials

| Service    | Credentials (username/password) |
| ---------- | ------------------------------- |
| PostgreSQL | postgres/example                |
