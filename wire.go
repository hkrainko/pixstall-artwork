//+build wireinject

package main

import (
	"github.com/google/wire"
	"github.com/streadway/amqp"
	"go.mongodb.org/mongo-driver/mongo"
	"pixstall-artwork/app/artwork/delivery/http"
	artwork_repo "pixstall-artwork/app/artwork/repo/mongo"
	artwork_ucase "pixstall-artwork/app/artwork/usecase"
	msg_broker_repo "pixstall-artwork/app/msg-broker/repo/rabbitmq"
)

func InitArtworkController(db *mongo.Database, conn *amqp.Connection) http.ArtworkController {
	wire.Build(
		http.NewArtworkController,
		artwork_ucase.NewArtworkUseCase,
		artwork_repo.NewMongoArtworkRepo,
		msg_broker_repo.NewRabbitMQMsgBrokerRepo,
	)

	return http.ArtworkController{}
}

func InitCommissionMessageBroker(db *mongo.Database)
