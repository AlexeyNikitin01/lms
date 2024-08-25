package server

import (
	amqp "github.com/rabbitmq/amqp091-go"
)

type Server struct {
	Conn *amqp.Connection
}

func NewServer(conn *amqp.Connection) *Server {
	return &Server{
		Conn: conn,
	}
}
