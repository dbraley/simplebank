# TODO: make help
# TODO: add some checks and dependencies so running the command multiple times doesn't fail and you only need to run `make dbup` etc
# TODO: try using a make alternative for this

postgres:
	docker run --name postgres13 -p 5432:5432 -e POSTGRES_USER=postgres -e POSTGRES_PASSWORD=secret -d postgres:13-alpine

createdb:
	docker exec -it postgres13 createdb --username=postgres --owner=postgres simple_bank

dropdb:
	docker exec -it postgres13 dropdb simple_bank

migrateup:
	migrate -path db/migration -database "postgresql://localhost:5432/postgres?sslmode=disable" -verbose up

migratedown:
	migrate -path db/migration -database "postgresql://localhost:5432/postgres?sslmode=disable" -verbose down

sqlc:
	sqlc generate

test:
	go test -v -cover ./...

PHONY: createdb dropdb postgres