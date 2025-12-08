package notification

import (
	"cinema_service/internal/domain"
	"encoding/json"
	"fmt"

	amqp "github.com/rabbitmq/amqp091-go"
)

type amqpNotificationSender struct {
	conn      *amqp.Connection
	channel   *amqp.Channel
	queueName string
}

func NewAMQPNotificationSender(amqpURL, queueName string) (*amqpNotificationSender, error) {
	conn, err := amqp.Dial(amqpURL)
	if err != nil {
		return nil, fmt.Errorf("amqp dial error: %w", err)
	}

	ch, err := conn.Channel()
	if err != nil {
		_ = conn.Close()
		return nil, fmt.Errorf("amqp channel error: %w", err)
	}

	// гарантируем, что очередь существует
	_, err = ch.QueueDeclare(
		queueName,
		true,  // durable
		false, // autoDelete
		false, // exclusive
		false, // noWait
		nil,   // args
	)
	if err != nil {
		_ = ch.Close()
		_ = conn.Close()
		return nil, fmt.Errorf("queue declare error: %w", err)
	}

	return &amqpNotificationSender{
		conn:      conn,
		channel:   ch,
		queueName: queueName,
	}, nil
}

// Желательно вызвать при завершении сервиса, но для диплома можно и не дергать
func (s *amqpNotificationSender) Close() {
	if s.channel != nil {
		_ = s.channel.Close()
	}
	if s.conn != nil {
		_ = s.conn.Close()
	}
}

type emailMessage struct {
	Recipient string `json:"recipient"`
	Subject   string `json:"subject"`
	Body      string `json:"body"`
}

func (s *amqpNotificationSender) SendTicketBoughtNotification(t domain.Ticket) error {
	msg := emailMessage{
		Recipient: t.Email,
		Subject:   "Покупка билета в кино",
		Body: fmt.Sprintf(
			"Вы купили билет на сеанс %d.\nРяд: %d, место: %d.",
			t.SessionID, t.Row, t.Seat,
		),
	}

	body, err := json.Marshal(msg)
	if err != nil {
		return fmt.Errorf("marshal notification: %w", err)
	}

	err = s.channel.Publish(
		"",          // default exchange
		s.queueName, // routing key = имя очереди (ticket-queue)
		false,
		false,
		amqp.Publishing{
			ContentType: "application/json",
			Body:        body,
		},
	)
	if err != nil {
		return fmt.Errorf("publish: %w", err)
	}

	return nil
}
