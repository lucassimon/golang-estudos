PWD := $(shell pwd)
DB_USER := estudos
DB_PASSWORD := teste123
DB_NAME := golangestudos
DB_PORT := 25432

start_postgres:
	docker start postgis
	docker exec postgis service postgresql restart

migrateup:
	migrate -path $(PWD)/databases/migrations -database "postgresql://$(DB_USER):$(DB_PASSWORD)@localhost:$(DB_PORT)/$(DB_NAME)?sslmode=disable" -verbose up

migratedown:
	migrate -path $(PWD)/databases/migrations -database "postgresql://$(DB_USER):$(DB_PASSWORD)@localhost:$(DB_PORT)/$(DB_NAME)?sslmode=disable" -verbose down


migrateupto:
	migrate -path $(PWD)/databases/migrations -database "postgresql://$(DB_USER):$(DB_PASSWORD)@localhost:$(DB_PORT)/$(DB_NAME)?sslmode=disable" -verbose up $(VERSION)


dropdb:
    docker exec postgis dropdb $(DB_NAME)



.PHONY: start_postgres
