postgres:
	docker run --name postgres15 -p 5432:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=secret postgres:15-alpine
migrateup:
	migrate -path db/migration -database "postgresql://root:secret@localhost:5432/simple_bank?sslmode=disable"	 -verbose up
migratedown:
	migrate -path db/migration -database "postgresql://root:secret@localhost:5432/simple_bank?sslmode=disable"	 -verbose down

createdb:
	docker exec -it postgres15 createdb --username=root --owner=root simple_bank
dropdb:
	docker exec -it postgres15 dropdb simple_bank

sqlc:
	sqlc generate

.PHONY: postgres dropdb createdb  migratedown  migrateup
