package app

import (
	"github.com/Video-Quality-Enhancement/VQE-Response-Producer/internal/consumers"
	"github.com/Video-Quality-Enhancement/VQE-Response-Producer/internal/producers"
	"github.com/Video-Quality-Enhancement/VQE-Response-Producer/internal/repositories"
	"github.com/Video-Quality-Enhancement/VQE-Response-Producer/internal/services"
	amqp "github.com/rabbitmq/amqp091-go"
	"go.mongodb.org/mongo-driver/mongo"
)

func SetUpEnhancedVideoConsumer(collection *mongo.Collection, consumerCh, notifyProducerCh *amqp.Channel) consumers.EnhancedVideoConsumer {

	respository := repositories.NewEnhancedVideoRepository(collection)
	producer := producers.NewNotificationProducer(notifyProducerCh)
	service := services.NewEnhancedVideoService(respository, producer)
	consumer := consumers.NewEnhancedVideoConsumer(consumerCh, service)
	return consumer

}
