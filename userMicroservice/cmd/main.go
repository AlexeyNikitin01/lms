package main

import (
	"fmt"
	"log"

	"edu-material/userMicroservice/internal/app"
	"edu-material/userMicroservice/internal/ports/httpgin"
)

func main() {
	fmt.Println("init server")

	a := app.CreateAppUser()

	s := httpgin.Server(":18080", a)

	log.Fatal(s.ListenAndServe())
}
