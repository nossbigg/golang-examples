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

db-new-migration:
	migrate create -ext sql -dir migrations $(file)

db-up:
	migrate -database "postgres://postgres:example@localhost:5432/golangexamples?sslmode=disable" -path migrations up 1

db-down:
	migrate -database "postgres://postgres:example@localhost:5432/golangexamples?sslmode=disable" -path migrations down 1