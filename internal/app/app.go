package app

import (
	"os"

	"github.com/Video-Quality-Enhancement/VQE-Enhanced-Video-Response/internal/config"
	"go.mongodb.org/mongo-driver/mongo"
)

func SetUpApp(database *mongo.Database, conn config.AMQPconnection, firebaseClient config.FirebaseClient) {

	videoCollection := database.Collection(os.Getenv("VIDEO_ENHANCE_COLLECTION"))
	userCollection := database.Collection(os.Getenv("USER_COLLECTION"))
	consumer := SetUpEnhancedVideoConsumer(videoCollection, userCollection, firebaseClient, conn)
	consumer.Consume()

}
