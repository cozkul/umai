backend_container_id := $(shell docker ps -aqf name=backend -n 1)
frontend_container_id := $(shell docker ps -aqf name=frontend -n 1)
db_conatiner_id := $(shell docker ps -aqf name=db -n 1)

db_user := $(shell grep "DB_USER=" .env | cut -d = -f 2)
db_password := $(shell grep "DB_PASSWORD=" .env | cut -d = -f 2)
db_name := $(shell grep "DB_NAME=" .env | cut -d = -f 2)

psql:
	docker exec -it $(db_conatiner_id) psql -U $(db_user)

backend-sh:
	docker exec -it $(backend_container_id) sh

frontend-sh:
	docker exec -it $(frontend_container_id) sh

createdb:
	docker exec -it $(db_conatiner_id) createdb --username=umai --owner=$(db_user) $(db_name)

dropdb:
	docker exec -it $(db_conatiner_id) dropdb -U $(db_user) $(db_name)

migrateup:
	migrate -path server/database/migrations -database "postgresql://$(db_user):$(db_password)@localhost:5432/$(db_name)?sslmode=disable" -verbose up

migratedown:
	migrate -path server/database/migrations -database "postgresql://$(db_user):$(db_password)@localhost:5432/$(db_name)?sslmode=disable" -verbose down

.PHONY: postgresinit postgres createdb dropdb migrateup migratedown