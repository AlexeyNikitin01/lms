user-postgres:
	docker run --name user -e POSTGRES_PASSWORD=pass -e POSTGRES_USER=postgres -e POSTGRES_DB=user -p 7777:5432 -d postgres

db-run:
	docker start user

db-stop:
	docker stop user

sql-gen:
	sqlboiler -c etc/config.template.yml -p entity -o internal/repository/pg/entity --struct-tag-casing camel --add-soft-deletes --no-tests --wipe psql

grpc-user:
	protoc --go_out=. --go-grpc_out=. api/*.proto

pg-migrate:
	migrate -path ./internal/migrations/ -database postgres://postgres:pass@localhost:7777/user?sslmode=disable up

run:
	go run ./cmd/main.go