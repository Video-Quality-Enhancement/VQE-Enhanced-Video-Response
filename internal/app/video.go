package app

import (
	"github.com/Video-Quality-Enhancement/VQE-Enhanced-Video-Response/internal/config"
	"github.com/Video-Quality-Enhancement/VQE-Enhanced-Video-Response/internal/consumers"
	"github.com/Video-Quality-Enhancement/VQE-Enhanced-Video-Response/internal/producers"
	"github.com/Video-Quality-Enhancement/VQE-Enhanced-Video-Response/internal/repositories"
	"github.com/Video-Quality-Enhancement/VQE-Enhanced-Video-Response/internal/services"
	"go.mongodb.org/mongo-driver/mongo"
)

func SetUpEnhancedVideoConsumer(videoCollection, userCollection *mongo.Collection, firebaseClient config.FirebaseClient, conn config.AMQPconnection) consumers.EnhancedVideoConsumer {

	videoRepository := repositories.NewEnhancedVideoRepository(videoCollection)
	notifyProducer := producers.NewNotificationProducer(conn)

	userRepository := repositories.NewUserRepository(userCollection)
	userService := services.NewUserService(userRepository, firebaseClient)

	videoService := services.NewEnhancedVideoService(videoRepository, userService, notifyProducer)
	consumer := consumers.NewEnhancedVideoConsumer(conn, videoService)
	return consumer

}
