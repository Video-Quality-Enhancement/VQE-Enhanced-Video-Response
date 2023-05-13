package video

import (
	"github.com/Video-Quality-Enhancement/VQE-Response-Producer/internal/config"
	"github.com/Video-Quality-Enhancement/VQE-Response-Producer/internal/video/consumers"
	"github.com/Video-Quality-Enhancement/VQE-Response-Producer/internal/video/repositories"
	"github.com/Video-Quality-Enhancement/VQE-Response-Producer/internal/video/services"
	"go.mongodb.org/mongo-driver/mongo"
)

func SetUpConsumer(collection *mongo.Collection, conn config.AMQPconnection) consumers.EnhancedVideoConsumer {

	respository := repositories.NewEnhancedVideoRepository(collection)
	service := services.NewEnhancedVideoService(respository)
	consumer := consumers.NewEnhancedVideoConsumer(conn, service)
	return consumer

}
