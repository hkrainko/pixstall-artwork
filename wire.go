//+build wireinject

package main

import (
	"github.com/google/wire"
	"github.com/streadway/amqp"
	"go.mongodb.org/mongo-driver/mongo"
	"pixstall-artwork/app/artwork/delivery/http"
	artwork_repo "pixstall-artwork/app/artwork/repo/mongo"
	artwork_ucase "pixstall-artwork/app/artwork/usecase"
	"pixstall-artwork/app/commission/delivery/rabbitmq"
	msg_broker_repo "pixstall-artwork/app/msg-broker/repo/rabbitmq"
	rabbitmq2 "pixstall-artwork/app/user/delivery/rabbitmq"
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

func InitCommissionMessageBroker(db *mongo.Database, conn *amqp.Connection) rabbitmq.CommissionMessageBroker {
	wire.Build(
		rabbitmq.NewCommissionMessageBroker,
		artwork_ucase.NewArtworkUseCase,
		artwork_repo.NewMongoArtworkRepo,
		msg_broker_repo.NewRabbitMQMsgBrokerRepo,
		)
	return rabbitmq.CommissionMessageBroker{}
}

func InitUserMessageBroker(db *mongo.Database, conn *amqp.Connection) rabbitmq2.UserMessageBroker {
	wire.Build(
		rabbitmq2.NewUserMessageBroker,
		artwork_ucase.NewArtworkUseCase,
		artwork_repo.NewMongoArtworkRepo,
		msg_broker_repo.NewRabbitMQMsgBrokerRepo,
	)
	return rabbitmq2.UserMessageBroker{}
}