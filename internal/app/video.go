package app

import (
	"github.com/Video-Quality-Enhancement/VQE-Response-Producer/internal/consumers"
	"github.com/Video-Quality-Enhancement/VQE-Response-Producer/internal/repositories"
	"github.com/Video-Quality-Enhancement/VQE-Response-Producer/internal/services"
	amqp "github.com/rabbitmq/amqp091-go"
	"go.mongodb.org/mongo-driver/mongo"
)

func SetUpEnhancedVideoConsumer(collection *mongo.Collection, consumerCh, notifyProducerCh *amqp.Channel) consumers.EnhancedVideoConsumer {

	respository := repositories.NewEnhancedVideoRepository(collection)
	service := services.NewEnhancedVideoService(respository, notifyProducerCh)
	consumer := consumers.NewEnhancedVideoConsumer(consumerCh, service)
	return consumer

}
