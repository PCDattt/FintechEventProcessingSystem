package rabbitmq

import (
	"encoding/json"

	"github.com/PCDattt/FintechEventProcessingSystem/shared/model"
	"github.com/streadway/amqp"
)

type Publisher struct {
	conn *amqp.Connection
	channel *amqp.Channel
	queue amqp.Queue
}

func NewPublisher(rabbitURL string, queueName string) (*Publisher, error) {
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

	return &Publisher {
		conn: conn,
		channel: ch,
		queue: q,
	}, nil
}

func (p *Publisher) PublishTransaction(tx model.Transaction) error {
	body, err := json.Marshal(tx)
	if err != nil {
		return err
	}

	return p.channel.Publish(
		"",
		p.queue.Name,
		false,
		false,
		amqp.Publishing {
			ContentType: "application/json",
			Body: body,
		},
	)
}

func (p *Publisher) Close() {
	p.channel.Close()
	p.conn.Close()
}