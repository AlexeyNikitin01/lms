user-postgres:
	docker run --name user -e POSTGRES_PASSWORD=pass -e POSTGRES_USER=postgres -e POSTGRES_DB=user -p 7777:5432 -d postgres

postgres-start:
	docker start user

postgres-stop:
	docker stop user

sql-gen:
	sqlboiler -c etc/config.template.yml -p entity -o internal/repository/pg/entity --struct-tag-casing camel --add-soft-deletes --no-tests --wipe psql

grpc-user:
	protoc --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative internal/ports/grpc/user.proto

pg-migrate:
	migrate -path ./internal/migrations/ -database postgres://postgres:pass@localhost:7777/user?sslmode=disable up
