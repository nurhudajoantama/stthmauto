package hmstt

import (
	"context"
	"time"

	"github.com/nurhudajoantama/stthmauto/internal/rabbitmq"
	amqp "github.com/rabbitmq/amqp091-go"
	"github.com/rs/zerolog/log"
)

type hmsttEvent struct {
	ch *amqp.Channel
	q  *amqp.Queue
}

func NewEvent(conn *amqp.Connection) *hmsttEvent {

	ch := rabbitmq.NewRabbitMQChannel(conn)

	q, err := ch.QueueDeclare(
		MQ_CHANNEL_HMSTT, // name
		false,            // durable
		true,             // delete when unused
		false,            // exclusive
		false,            // no-wait
		nil,              // arguments
	)
	if err != nil {
		log.Error().Err(err).Msg("Failed to declare a queue")
		return nil
	}

	return &hmsttEvent{
		ch: ch,
		q:  &q,
	}
}

func (e *hmsttEvent) StateChange(ctx context.Context, key string, value string) error {
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	err := e.ch.PublishWithContext(
		ctx,
		"",       // exchange
		e.q.Name, // routing key
		false,    // mandatory
		true,     // immediate
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte(value),
		},
	)

	return err
}
