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

func InitSubscriber(topic string) SubscriberInterface {
	logger := log.InitLogger()
	cfg := config.LoadRabbitMQConfig()

	if strings.TrimSpace(topic) == "" {
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
		topic,
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

	q, err := ch.QueueDeclare(
		"",
		false,
		false,
		true,
		false,
		nil,
	)
	if err != nil {
		logger.Fatalf("Failed to declare an queue: %v", err)
	}

	err = ch.QueueBind(
		q.Name,
		"",
		topic,
		false,
		nil,
	)
	if err != nil {
		logger.Fatalf("Failed to bind a queue: %v", err)
	}

	logger.Infof("RabbitMQ subscriber listening on %s", topic)

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
			s.logger.Infof("EVENT: %s", d.Body)
			fn(c, NewEvent(c, d.Body))
		}
	}()

	<-forever
	return nil
}
