DB_URL=postgresql://root:pickaxe-db@localhost:5432/pickaxe_db?sslmode=disable

sqlc:
	sqlc generate

db_docs:
	dbdocs build db/docs/pickaxe.dbml

db_schema:
	dbml2sql --postgres -o db/docs/schema.sql db/docs/pickaxe.dbml

postgres:
	docker run --name pickaxe -p 5432:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=pickaxe-db -d postgres:15-alpine

docker-network:
	docker network create pickaxe-network

postgres-network:
	docker run --name pickaxe --network pickaxe-network -p 5432:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=pickaxe-db -d postgres:alpine3.14

createdb:
	docker exec -it pickaxe createdb --username=root --owner=root pickaxe_db

dropdb:
	docker exec -it pickaxe dropdb pickaxe_db

migrateup:
	migrate -path db/migration -database "$(DB_URL)" -verbose up

migratedown:
	migrate -path db/migration -database "$(DB_URL)" -verbose down

pickaxe:
	go run cmd/pickaxe/main.go

psocket:
	go run cmd/psocket/main.go 

build:
	go build -o bin/pickaxe -v cmd/pickaxe/main.go
	go build -o bin/psocket -v cmd/psocket/main.go

install-go:
	go install -v ./...

docker-build-pickaxe:
	docker build -f Dockerfile.pickaxe -t pickaxe:latest .      

docker-build-psocket:
	docker build -f Dockerfile.psocket -t psocket:latest .                                                                

docker-container-pickaxe:
	docker run --name pickaxe_app --network pickaxe-network -p 8080:8080 -e GIN_MODE=release -e SOCKET_ADDRESS=psocket:8081 -e DB_SOURCE="postgresql://root:pickaxe-db@pickaxe:5432/pickaxe_db?sslmode=disable" pickaxe:latest

docker-container-psocket:
	docker run --name psocket --network pickaxe-network psocket:latest

docker-compose:
	docker compose up

.PHONY: sqlc db_docs db_schema postgres docker-network postgres-network createdb migrateup migratedown build install-go docker-build-pickaxe docker-build-psocket docker-container-pickaxe docker-container-psocket docker-compose go psocket pickaxe