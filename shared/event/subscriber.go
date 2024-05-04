package event

import (
	"context"
	"fmt"
	"go-boilerplate/config"
	"go-boilerplate/shared/errors"
	"go-boilerplate/shared/log"
	"strings"

	amqp "github.com/rabbitmq/amqp091-go"
	"go.uber.org/zap"
)

type Subscriber struct {
	channel *amqp.Channel
	logger  *zap.SugaredLogger
	queue   string
}

type SubscriberOptions struct {
	Topic string
	Queue *string
}

func InitSubscriber(options SubscriberOptions) SubscriberInterface {
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
		logger.Fatal(errors.NewInitializationError(err, "rabbitMQ subscriber").Message)
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

	var qName string = fmt.Sprint("queue-", options.Topic)
	if options.Queue != nil {
		qName = *options.Queue
	}

	q, err := ch.QueueDeclare(
		qName,
		false,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		logger.Fatalf("Failed to declare an queue: %v", err)
	}

	err = ch.QueueBind(
		q.Name,
		"",
		options.Topic,
		false,
		nil,
	)
	if err != nil {
		logger.Fatalf("Failed to bind a queue: %v", err)
	}

	logger.Infof("RabbitMQ subscriber listening on %s", options.Topic)

	return Subscriber{
		channel: ch,
		logger:  logger,
		queue:   q.Name,
	}
}

func (s Subscriber) Subscribe(c context.Context, fn func(c context.Context, event Event) error) error {
	msgs, err := s.channel.Consume(
		s.queue,
		"",
		true,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		s.logger.Fatalf("Failed to register a consumer", err)
		return err
	}

	var forever chan struct{}

	go func() {
		for d := range msgs {
			event, err := Deserialize(d.Body)
			if err != nil {
				s.logger.Fatalf("Failed to serialize event: %v", err)
			}

			fn(c, event)
		}
	}()

	<-forever
	return nil
}

func (s Subscriber) Shutdown() error {
	return s.channel.Close()
}
