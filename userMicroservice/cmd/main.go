package main

import (
	"fmt"
	"log"

	"edu-material/userMicroservice/internal/app"
	"edu-material/userMicroservice/internal/adapters/postgres"
	"edu-material/userMicroservice/internal/ports/httpgin"
)

func main() {
	fmt.Println("init server")

	cfg := postgres.Config{}
	db, err := postgres.CreatePostgres(&cfg)
	if err != nil {
		log.Fatal(err)
	}
	
	a := app.CreateAppUser(postgres.CreateRepoUser(db))

	s := httpgin.Server(":18080", a)

	log.Fatal(s.ListenAndServe())
}
