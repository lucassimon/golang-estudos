PWD := $(shell pwd)
DB_USER := postgres
DB_PASSWORD := postgres
DB_NAME := golangestudos
DB_PORT := 25432

start_postgres:
	docker start postgis
	docker exec postgis service postgresql restart

migrateup:
	migrate -path $(PWD)/migrations -database "postgresql://$(DB_USER):$(DB_PASSWORD)@localhost:$(DB_PORT)/$(DB_NAME)?sslmode=disable" -verbose up

migratedown:
	migrate -path $(PWD)/migrations -database "postgresql://$(DB_USER):$(DB_PASSWORD)@localhost:$(DB_PORT)/$(DB_NAME)?sslmode=disable" -verbose down


migrateupto:
	migrate -path $(PWD)/migrations -database "postgresql://$(DB_USER):$(DB_PASSWORD)@localhost:$(DB_PORT)/$(DB_NAME)?sslmode=disable" -verbose up $(VERSION)


teste:



.PHONY: start_postgres
