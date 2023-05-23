package app

import (
	"os"

	"github.com/Video-Quality-Enhancement/VQE-Response-Producer/internal/config"
	"go.mongodb.org/mongo-driver/mongo"
)

func SetUpApp(database *mongo.Database, conn config.AMQPconnection) {

	collection := database.Collection(os.Getenv("VIDEO_COLLECTION"))
	consumerCh := conn.NewChannel()
	notifyProducerCh := conn.NewChannel()
	consumer := SetUpEnhancedVideoConsumer(collection, consumerCh, notifyProducerCh)
	consumer.Consume()

}
