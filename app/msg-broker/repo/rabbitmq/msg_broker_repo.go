package rabbitmq

import (
	"context"
	"encoding/json"
	"github.com/streadway/amqp"
	"log"
	"pixstall-artwork/app/msg-broker/repo/rabbitmq/msg"
	"pixstall-artwork/domain/artwork/model"
	msg_broker "pixstall-artwork/domain/msg-broker"
)

type rabbitmqMsgBrokerRepo struct {
	ch *amqp.Channel
}

func NewRabbitMQMsgBrokerRepo(conn *amqp.Connection) msg_broker.Repo {
	ch, err := conn.Channel()
	if err != nil {
		log.Fatalf("Failed to open a channel %v", err)
	}
	err = ch.Qos(5, 0, false)
	if err != nil {
		log.Fatalf("Failed to set QoS %v", err)
	}
	return rabbitmqMsgBrokerRepo{
		ch: ch,
	}
}

func (r rabbitmqMsgBrokerRepo) SendArtworkCreatedMessage(ctx context.Context, artwork model.Artwork) error {
	cArtwork := msg.NewCreatedArtwork(artwork)
	b, err := json.Marshal(cArtwork)
	if err != nil {
		return err
	}
	err = r.ch.Publish(
		"artwork",
		"artwork.event.created",
		false,
		false,
		amqp.Publishing{
			ContentType: "application/json",
			Body:        b,
		},
	)
	if err != nil {
		return err
	}
	return nil
}