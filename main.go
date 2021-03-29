package main

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/streadway/amqp"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"time"
)

func main() {

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)

	//Mongo
	dbClient, err := mongo.Connect(ctx, options.Client().ApplyURI("mongodb://localhost:27017"))
	if err != nil {
		panic(err)
	}
	defer cancel()
	defer func() {
		if err = dbClient.Disconnect(ctx); err != nil {
			panic(err)
		}
	}()
	_ = dbClient.Database("pixstall-artwork")

	//RabbitMQ
	rbMQConn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	if err != nil {
		log.Fatalf("Failed to connect to RabbitMQ %v", err)
	}
	defer rbMQConn.Close()
	ch, err := rbMQConn.Channel()
	if err != nil {
		log.Fatalf("Failed to create channel %v", err)
	}
	err = ch.ExchangeDeclare(
		"artwork",
		"topic",
		true,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		log.Fatalf("Failed to create exchange %v", err)
	}

	//commMsgBroker := InitCommissionMessageBroker(db, rbMQConn, awsS3, hub)
	//go commMsgBroker.StartUpdateCommissionQueue()
	//go commMsgBroker.StartCommissionValidatedQueue()
	//go commMsgBroker.StartCommissionMessageDeliverQueue()
	//defer commMsgBroker.StopAllQueues()

	// Gin
	r := gin.Default()

	//userIDExtractor := middleware.NewJWTPayloadsExtractor([]string{"userId"})

	//apiGroup := r.Group("/api")
	//artworkGroup := apiGroup.Group("/artworks")
	//{
	//	ctrl := InitArtworkController(db, rbMQConn)
	//	artworkGroup.GET("", userIDExtractor.ExtractPayloadsFromJWTInHeader, ctrl.)
	//	artworkGroup.GET("/:id", userIDExtractor.ExtractPayloadsFromJWTInHeader, ctrl.GetCommission)
	//	artworkGroup.GET("/:id/details", userIDExtractor.ExtractPayloadsFromJWTInHeader, ctrl.GetCommissionDetails)
	//	artworkGroup.GET("/:id/messages", userIDExtractor.ExtractPayloadsFromJWTInHeader, ctrl.GetMessages)
	//	artworkGroup.POST("/:id/messages", userIDExtractor.ExtractPayloadsFromJWTInHeader, ctrl.CreateMessage)
	//	artworkGroup.POST("", userIDExtractor.ExtractPayloadsFromJWTInHeader, ctrl.AddCommission)
	//	artworkGroup.PATCH("/:id", userIDExtractor.ExtractPayloadsFromJWTInHeader, ctrl.UpdateCommission)
	//}

	err = r.Run(":9005")
	print(err)
}