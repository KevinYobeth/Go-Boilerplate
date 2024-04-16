package event

import (
	"fmt"
	"go-boilerplate/config"
	"go-boilerplate/shared/errors"
	"go-boilerplate/shared/log"
	"go-boilerplate/shared/utils"
	"strings"

	amqp "github.com/rabbitmq/amqp091-go"
	"go.uber.org/zap"
)

type Publisher struct {
	channel *amqp.Channel
	log     *zap.SugaredLogger
	topic   string
}

func InitPublisher(topic string) PublisherInterface {
	log := log.InitLogger()
	cfg := config.LoadRabbitMQConfig()

	if strings.TrimSpace(topic) == "" {
		log.Fatal("Topic cannot be empty")
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
		log.Fatal(errors.NewInitializationError(err, "rabbitMQ publisher").Message)
	}

	ch, err := conn.Channel()
	if err != nil {
		log.Fatalf("Failed to open a channel: %v", err)
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
		log.Fatalf("Failed to declare an exchange: %v", err)
	}

	return Publisher{
		channel: ch,
		log:     log,
		topic:   topic,
	}
}

func (p Publisher) Publish(event Event) error {
	json, err := utils.ToJsonString(event.data)
	if err != nil {
		p.log.Fatalf("Failed to marshal event: %v", err)
		return err
	}

	err = p.channel.PublishWithContext(event.c,
		p.topic,
		"",
		false,
		false,
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte(json),
		})
	if err != nil {
		p.log.Fatalf("Failed to open a channel: %v", err)
		return err
	}

	return nil
}
