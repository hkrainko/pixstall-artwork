package rabbitmq

import (
	"context"
	"encoding/json"
	"github.com/streadway/amqp"
	"log"
	"pixstall-artwork/app/commission/delivery/rabbitmq/msg"
	"pixstall-artwork/domain/artwork"
	"time"
)

type CommissionMessageBroker struct {
	artworkUseCase artwork.UseCase
	ch             *amqp.Channel
}

func NewCommissionMessageBroker(artworkUseCase artwork.UseCase, conn *amqp.Connection) CommissionMessageBroker {
	ch, err := conn.Channel()
	if err != nil {
		log.Fatalf("Failed to open a channel %v", err)
	}
	err = ch.Qos(5, 0, false)
	if err != nil {
		log.Fatalf("Failed to set QoS %v", err)
	}

	return CommissionMessageBroker{
		artworkUseCase: artworkUseCase,
		ch:             ch,
	}
}

func (c CommissionMessageBroker) StartCompletionToArtworkQueue() {
	//TODO
	q, err := c.ch.QueueDeclare(
		"commission-completion-to-artwork", // name
		true,                               // durable
		false,                              // delete when unused
		false,                              // exclusive
		false,                              // no-wait
		nil,                                // arguments
	)
	if err != nil {
		log.Fatalf("Failed to declare a queue %v", err)
	}
	err = c.ch.QueueBind(
		q.Name,
		"commission.event.completed",
		"commission",
		false,
		nil,
	)
	if err != nil {
		log.Fatalf("Failed to bind queue %v", err)
	}

	msgs, err := c.ch.Consume(
		q.Name, // queue
		"",     // consumer
		false,  // auto-ack
		false,  // exclusive
		false,  // no-local
		false,  // no-wait
		nil,    // args
	)
	if err != nil {
		log.Fatalf("Failed to register a consumer %v", err)
	}

	forever := make(chan bool)

	go func() {
		for d := range msgs {
			log.Printf("Received a message: %s", d.Body)
			d.Ack(false)

			ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
			go func() {
				for {
					select {
					case <-ctx.Done():
						switch ctx.Err() {
						case context.DeadlineExceeded:
							log.Println("context.DeadlineExceeded")
						case context.Canceled:
							log.Println("context.Canceled")
						default:
							log.Println("default")
						}
						return // returning not to leak the goroutine
					}
				}
			}()

			switch d.RoutingKey {
			case "commission.event.completed":
				err := c.completionToArtwork(ctx, d.Body)
				if err != nil {
					//TODO: error handling, store it ?
				}
				cancel()
			default:
				cancel()
			}
		}
	}()

	<-forever
}

// Private handler
func (c CommissionMessageBroker) completionToArtwork(ctx context.Context, body []byte) error {
	req := msg.CompletedCommission{}
	err := json.Unmarshal(body, &req)
	if err != nil {
		return err
	}
	return c.artworkUseCase.AddArtwork(ctx, req.Commission)
}
