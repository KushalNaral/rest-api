package queue

import (
	"encoding/json"
	"errors"
	"log"
)

type EmailService interface {
	SendVerificationMail(email, token string) error
}

func StartEmailConsumer(rmq *RabbitMQ, emailService EmailService) error {

	msgs, err := rmq.ch.Consume(
		EmailVerification,
		"",
		false,
		false,
		false,
		false,
		nil)
	if err != nil {
		return errors.New("error starting EmailVerification consumption: " + err.Error())
	}

	go func() {
		for msg := range msgs {
			var payload EmailVerificationPayload
			if err := json.Unmarshal(msg.Body, &payload); err != nil {
				log.Println("error unmarshalling email verification payload: " + err.Error())
				msg.Nack(false, false) // no requeue if marshan fails
				continue
			}

			err := emailService.SendVerificationMail(payload.Email, payload.Token)
			if err != nil {
				log.Println("error sending email verification mail: " + err.Error())
				msg.Nack(false, true) // requeue here
			}

			log.Println("verification email sending success")
			msg.Ack(false) // success case
		}
	}()

	log.Println("email consumer started...")
	return nil

}
