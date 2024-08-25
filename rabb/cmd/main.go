package main

import (
	"fmt"
	"log"

	amqp "github.com/rabbitmq/amqp091-go"

	"rabb/pkg/server"
)

func main() {
	go func() {
		port := "amqp://guest:guest@localhost:5672/"

		var conn *amqp.Connection
		var err error

		for {
			conn, err = amqp.Dial(port)
			if err != nil {
				continue
			}
			break
		}

		defer conn.Close()

		fmt.Println("c", conn)
		log.Println("Connected to AMQP")
		s := server.NewServer(conn)
		log.Println("Read")
		err = s.ReadHelloWorld()
		if err != nil {
			log.Fatal(err)
		}
	}()

	fmt.Println("it`s work")

	f := make(chan struct{})
	<-f
}

//ctx := context.Background()
//

//log.Println("Send")
//err = s.SendHelloWorld(ctx)
//if err != nil {
//log.Fatal(err)
//}
//
