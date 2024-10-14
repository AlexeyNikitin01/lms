postgres-up:
	docker run --name diplom -e POSTGRES_PASSWORD=pass -e POSTGRES_USER=postgres -e POSTGRES_DB=diplom -p 5432:5432 -d postgres

postgres-start:
	docker start diplom

postgres-stop:
	docker stop diplom

grafana:
	docker run -d -p 3000:3000 --name=grafana grafana/grafana

prometheus:
	docker run -d \
      -p 9090:9090 \
      -v ./etc/prometheus.yml \
      --name=prometheus \
      prom/prometheus