package server

import (
	"context"
	"log"

	"github.com/pkg/errors"
	amqp "github.com/rabbitmq/amqp091-go"
)

func (s Server) SendHelloWorld(ctx context.Context) error {
	ch, err := s.Conn.Channel()
	if err != nil {
		return errors.Wrap(err, "failed to open channel")
	}

	defer ch.Close()

	q, err := ch.QueueDeclare(
		"hello", // name
		false,   // durable
		false,   // delete when unused
		false,   // exclusive
		false,   // no-wait
		nil,     // arguments
	)
	if err != nil {
		return errors.Wrap(err, "Failed to declare a queue")
	}

	body := "Hello World!"
	err = ch.PublishWithContext(ctx,
		"",     // exchange
		q.Name, // routing key
		false,  // mandatory
		false,  // immediate
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte(body),
		})
	if err != nil {
		return errors.Wrap(err, "Failed to publish a message")
	}
	log.Printf(" [x] Sent %s\n", body)

	return nil
}
