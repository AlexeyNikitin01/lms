package server

import (
	"log"

	"github.com/pkg/errors"
)

func (s Server) ReadHelloWorld() error {
	ch, err := s.Conn.Channel()
	if err != nil {
		return errors.Wrap(err, "failed to open channel")
	}

	defer ch.Close()

	msgs, err := ch.Consume(
		"hello", // name
		"",      // consumer
		true,    // auto-ack
		false,   // exclusive
		false,   // no-local
		false,   // no-wait
		nil,     // argsments
	)
	if err != nil {
		return errors.Wrap(err, "failed to declare queue")
	}

	forever := make(chan struct{})

	go func() {
		for d := range msgs {
			log.Printf("Received a message: %s", d.Body)
		}
		forever <- struct{}{}
	}()

	log.Printf(" [*] Waiting for messages. To exit press CTRL+C")
	<-forever

	return nil
}
