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
}

func NewEvent(conn *amqp.Connection) *hmsttEvent {
	ch := rabbitmq.NewRabbitMQChannel(conn)
	return &hmsttEvent{
		ch: ch,
	}
}

func (e *hmsttEvent) StateChange(ctx context.Context, key string, value string) error {
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	err := e.ch.PublishWithContext(
		ctx,
		"amq.topic",      // exchange
		MQ_CHANNEL_HMSTT, // routing key
		false,            // mandatory
		false,            // immediate
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte(value),
		},
	)
	if err != nil {
		log.Error().Err(err).Msg("Failed to publish a message")
	}

	return err
}
