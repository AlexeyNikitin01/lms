postgres-up:
	docker run --name diplom -e POSTGRES_PASSWORD=pass -e POSTGRES_USER=postgres -e POSTGRES_DB=diplom -p 5432:5432 -d postgres

postgres-start:
	docker start diplom

postgres-stop:
	docker stop diplom
