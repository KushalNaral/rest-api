package queue

import (
	"encoding/json"
	"errors"
	"time"

	amqp "github.com/rabbitmq/amqp091-go"
)

type RabbitMQ struct {
	conn *amqp.Connection
	ch   *amqp.Channel
}

type EmailVerificationPayload struct {
	UserID string `json:"user_id"`
	Email  string `json:"email"`
	Token  string `json:"token"`
}

const (
	EmailVerification = "email.verification"
	ExchangeNames     = "auth_events"
)

func NewRabbitMQ(url string) (*RabbitMQ, error) {

	// remember to close this connection after use !! so resource for Dial doenst leak
	conn, err := amqp.Dial(url)
	if err != nil {
		return nil, errors.New("error connecting to RabbitMQ: " + err.Error())
	}

	ch, err := conn.Channel()
	if err != nil {
		return nil, errors.New("error opening a new RabbitMQ channel: " + err.Error())
	}

	err = ch.ExchangeDeclare(ExchangeNames, "direct", true, false, false, false, nil)
	if err != nil {
		return nil, errors.New("error declaring a new RabbitMQ exchange: " + err.Error())
	}

	_, err = ch.QueueDeclare(
		EmailVerification, true, false, false, false, nil)
	if err != nil {
		return nil, errors.New("error declaring a new RabbitMQ QueueEmailVerification queue: " + err.Error())
	}

	err = ch.QueueBind(EmailVerification, EmailVerification, ExchangeNames, false, nil)
	if err != nil {
		return nil, errors.New("error binding the RabbitMQ QueueEmailVerification queue: " + err.Error())
	}

	return &RabbitMQ{
		conn: conn,
		ch:   ch,
	}, nil
}

func (r *RabbitMQ) PublishEmailVerification(payload EmailVerificationPayload) error {
	body, err := json.Marshal(payload)
	if err != nil {
		return errors.New("error marshalling email verification payload: " + err.Error())
	}

	return r.ch.Publish(
		ExchangeNames,
		EmailVerification,
		false,
		false,
		amqp.Publishing{
			ContentType: "application/json",
			Body:        body,
			Timestamp:   time.Now(),
		},
	)
}

func (r *RabbitMQ) CloseConnection() error {
	err := r.conn.Close()
	if err != nil {
		return errors.New("error closing RabbitMQ connection: " + err.Error())
	}
	err = r.ch.Close()
	if err != nil {
		return errors.New("error closing RabbitMQ channel: " + err.Error())
	}
	return nil
}
