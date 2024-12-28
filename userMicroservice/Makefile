user-postgres:
	docker run --name user -e POSTGRES_PASSWORD=pass -e POSTGRES_USER=postgres -e POSTGRES_DB=user -p 7777:5432 -d postgres

db-run:
	docker start user

db-stop:
	docker stop user

sql-gen:
	sqlboiler -c etc/config.template.yml -p entity -o internal/repository/pg/entity --struct-tag-casing snake --add-soft-deletes --no-tests --wipe psql

grpc-user:
	protoc --go_out=. --go-grpc_out=. api/*.proto

pg-migrate:
	migrate -path ./internal/migrations/ -database postgres://postgres:pass@localhost:7777/user?sslmode=disable up

run:
	go run ./cmd/main.go

lint-without-fix:
	golangci-lint --version
	golangci-lint run --allow-parallel-runners --build-tags=integration --config ./build/ci/.golangci.yml ./...

generate:
	rm -rf mocks/*
	mockgen -destination=mocks/postgres_mock.go -package=mocks -source=internal/adapters/postgres/interface.go
	mockgen -destination=mocks/cloud_mock.go -package=mocks -source=internal/adapters/storage/interface.go
	mockgen -destination=mocks/app_mock.go -package=mocks -source=internal/app/app.go
	mockgen -destination=mocks/tel_mock.go -package=mocks -source=internal/metrics/face.go

prometheus:
	docker run \
      -d \
      --name prometheus_lms \
      -p 9090:9090 \
      -v ./etc/prometheus.yml:/etc/prometheus/prometheus.yml \
      prom/prometheus

run-prometheus:
	docker start prometheus

grafana:
	docker run -d \
	  --name=grafana \
	  -p 3000:3000 \
	  -e "GF_SECURITY_ADMIN_PASSWORD=admin" \
	  -v grafana-storage:/var/lib/grafana \
	  grafana/grafana:latest


run-grafana:
	docker start grafana

docker-build:
	docker build -t user .

docker-run:
	docker run -p 18080:18080 -p 50054:50054 user

all:
	make db-run
	make run
	