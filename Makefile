DB_URL=postgres://root:secret@localhost:5432/simple_bank?sslmode=disable

newnetwork:
	docker network create bank-network

postgres:
	docker run --name postgres12 --network bank-network -p 5432:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=secret -d postgres:12-alpine

createdb:
	docker exec -it postgres12 createdb --username=root --owner=root simple_bank

dropdb:
	docker exec -it postgres12 dropdb simple_bank

migrateinit:
	migrate create -ext sql -dir db/migration -seq init_schema

migrateup:
	migrate -path db/migration -database "$(DB_URL)" -verbose up

migratedown:
	migrate -path db/migration -database "$(DB_URL)" -verbose down

migrateup1:
	migrate -path db/migration -database "$(DB_URL)" -verbose up 1

migratedown1:
	migrate -path db/migration -database "$(DB_URL)" -verbose down 1

dbdocs:
	dbdocs build doc/db.dbml

dbschema:
	dbml2sql --postgres -o doc/schema.sql doc/db.dbml

sqlcwin:
	docker run --rm -v $(CURDIR):/src -w /src sqlc/sqlc generate

sqlcunix:
	sqlc generate

test:
	go test -v -cover ./...

server:
	go run main.go

mock:
	mockgen -destination db/mock/store.go  github.com/LKarrie/learn-go-project/db/sqlc Store

build:
	docker build -t simplebank:latest .

run:
	docker run --name simplebank --network bank-network -p 8080:8080 -e GIN_MODE=release -e DB_SOURCE="postgres://root:secret@postgres12:5432/simple_bank?sslmode=disable&TimeZone=Asia/Shanghai" simplebank:latest

proto:
	del .\pb\*.go
	protoc --proto_path=proto --go_out=pb --go_opt=paths=source_relative \
	--go-grpc_out=pb --go-grpc_opt=paths=source_relative \
	proto/*.proto

.PHONY: newnetwork postgres createdb dropdb migrateinit migrateup migratedown sqlcwin sqlcunix test server mock build run dbdocs dbschema proto