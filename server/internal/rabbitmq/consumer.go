package rabbitmq

import (
	"context"
	"encoding/json"
	"log"

	"github.com/PCDattt/FintechEventProcessingSystem/shared/model"
	"github.com/streadway/amqp"
)

type Consumer struct {
	conn *amqp.Connection
	channel *amqp.Channel
	queue amqp.Queue
}

func NewConsumer(rabbitURL string, queueName string) (*Consumer, error) {
	conn, err := amqp.Dial(rabbitURL)
	if err != nil {
		return nil, err
	}

	ch, err := conn.Channel()
	if err != nil {
		return nil, err
	}

	q, err := ch.QueueDeclare(
		queueName,
		true,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		return nil, err
	}

	return &Consumer{
		conn: conn,
		channel: ch,
		queue: q,
	}, nil
}

func (c *Consumer) StartConsuming(ctx context.Context, handler func(ctx context.Context, tx model.Transaction) (model.Transaction, error)) error {
	msgs, err := c.channel.Consume(
		c.queue.Name,
		"",
		true,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		return err
	}

	go func() {
		for msg := range msgs {
			var tx model.Transaction
			if err := json.Unmarshal(msg.Body, &tx); err != nil {
				log.Printf("Failed to unmarshal message: %v", err)
				continue
			}
			log.Printf("Received transaction: %v", tx)

			result, err := handler(ctx, tx)
			if err != nil {
				log.Printf("Transaction handler error: %v", err)
			}
			log.Printf("Processed transaction: %v", result)
		}
	}()

	return nil
}

func (c *Consumer) Close() {
	c.channel.Close()
	c.conn.Close()
}