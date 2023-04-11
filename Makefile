createdb:
	docker exec -it postgres15 createdb --username=root --owner=root shop_api

dropdb:
	docker exec -it postgres15 dropdb shop_api

migrateup:
	migrate -path db/migration -database "postgresql://root:secret@localhost:5432/shop_api?sslmode=disable" -verbose up

migratedown:
	migrate -path db/migration -database "postgresql://root:secret@localhost:5432/shop_api?sslmode=disable" -verbose down

.PHONY: 
	createdb, dropdb, migrateup, migratedown