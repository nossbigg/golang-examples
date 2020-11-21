# golang
go-test:
	go test ./...

go-fmt:
	gofmt -w .

# docker
docker-start:
	docker-compose up -d

docker-stop:
	docker-compose stop

docker-destroy:
	docker-compose down

docker-destroy-db:
	docker rm golangexamples_postgres_1

# db
db-makedb:
	createdb -h localhost -p 5432 -U postgres golangexamples

db-dropdb:
	dropdb -h localhost -p 5432 -U postgres golangexamples
