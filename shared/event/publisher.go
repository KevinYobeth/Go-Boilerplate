package event

import (
	"context"
	"fmt"
	"strings"

	"github.com/kevinyobeth/go-boilerplate/config"
	"github.com/kevinyobeth/go-boilerplate/shared/errors"
	"github.com/kevinyobeth/go-boilerplate/shared/log"

	amqp "github.com/rabbitmq/amqp091-go"
	"go.uber.org/zap"
)

type Publisher struct {
	channel *amqp.Channel
	logger  *zap.SugaredLogger
	topic   string
}

type PublisherOptions struct {
	Topic string
	Queue *string
}

func InitPublisher(options PublisherOptions) PublisherInterface {
	logger := log.InitLogger()
	cfg := config.LoadRabbitMQConfig()

	if strings.TrimSpace(options.Topic) == "" {
		logger.Fatal("Topic cannot be empty")
	}

	connStr := fmt.Sprintf(
		"amqp://%s:%s@%s:%s",
		cfg.RabbitMQUsername,
		cfg.RabbitMQPassword,
		cfg.RabbitMQHost,
		cfg.RabbitMQPort,
	)
	conn, err := amqp.Dial(connStr)
	if err != nil {
		logger.Fatal(errors.NewInitializationError(err, "rabbitMQ publisher").Message)
	}

	ch, err := conn.Channel()
	if err != nil {
		logger.Fatalf("Failed to open a channel: %v", err)
	}

	err = ch.ExchangeDeclare(
		options.Topic,
		"fanout",
		true,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		logger.Fatalf("Failed to declare an exchange: %v", err)
	}

	return Publisher{
		channel: ch,
		logger:  logger,
		topic:   options.Topic,
	}
}

func (p Publisher) Publish(c context.Context, event Event) error {
	body, err := Serialize(event)
	if err != nil {
		p.logger.Fatalf("Failed to serialize event: %v", err)
	}

	err = p.channel.PublishWithContext(c,
		p.topic,
		"",
		false,
		false,
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte(body),
		})
	if err != nil {
		p.logger.Fatalf("Failed to open a channel: %v", err)
		return err
	}

	return nil
}
