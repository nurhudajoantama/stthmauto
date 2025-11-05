package rabbitmq

import (
	"log"

	"github.com/nurhudajoantama/stthmauto/internal/config"
	amqp "github.com/rabbitmq/amqp091-go"
)

func NewRabbitMQConn(c config.MQTT) *amqp.Connection {
	conn, err := amqp.Dial(c.BrokerURL())
	if err != nil {
		log.Panicf("Failed to connect to RabbitMQ: %s", err)
	}

	return conn
}

func NewRabbitMQChannel(conn *amqp.Connection) *amqp.Channel {
	ch, err := conn.Channel()
	if err != nil {
		log.Panicf("Failed to open a channel: %s", err)
	}

	return ch
}
