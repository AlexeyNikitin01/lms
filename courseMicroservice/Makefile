DB_PORT=6666
DB_PASS=pass
DB_NAME=course
DB_USER=postgres
DB_HOST=localhost

PATH_MIGRATIONS = ./internal/migrations
PATH_API_PROTO = ./api/*.proto

DOCKER_NAME=courses

PATH_ENTITY_BOIL=internal/repository/pg/entity
PATH_CONF_BOIL=etc/config.template.yml

db-create:
	docker run --name $(DOCKER_NAME) -e POSTGRES_PASSWORD=$(DB_PASS) \
    -e POSTGRES_USER=$(DB_USER) \
    -e POSTGRES_DB=$(DB_NAME) \
    -p $(DB_PORT):5432 -d postgres

db-run:
	docker start $(DOCKER_NAME)

db-stop:
	docker stop $(DOCKER_NAME)

pg-migrate:
	migrate -path $(PATH_MIGRATIONS) -database postgres://$(DB_USER):$(DB_PASS)@$(DB_HOST):$(DB_PORT)/$(DB_NAME)?sslmode=disable up

grpc:
	protoc --go_out=. --go-grpc_out=. $(PATH_API_PROTO)

sql-gen:
	sqlboiler -c  $(PATH_CONF_BOIL) \
    -p entity \
    -o $(PATH_ENTITY_BOIL) \
    --struct-tag-casing snake \
    --add-soft-deletes \
    --no-tests \
    --wipe psql

run:
	go run ./cmd/main.go

mongo:
	docker run -d --name course_mongo -p 27017:27017 \
	-e MONGO_INITDB_ROOT_USERNAME=mongo \
	-e MONGO_INITDB_ROOT_PASSWORD=pass mongo
