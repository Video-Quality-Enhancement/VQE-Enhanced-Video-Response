package main

import (
	"os"

	"github.com/Video-Quality-Enhancement/VQE-Response-Producer/internal/config"
	"github.com/Video-Quality-Enhancement/VQE-Response-Producer/internal/video"
)

func init() {
	config.LoadEnvVariables()
}

func main() {

	logFile := config.SetupSlogOutputFile()
	defer logFile.Close()

	client := config.NewMongoClient()
	db := client.ConnectToDB()
	defer client.Disconnect()

	conn := config.NewAMQPconnection()
	defer conn.Disconnect()

	consumer := video.SetUpConsumer(db.Collection(os.Getenv("VIDEO_COLLECTION")), conn)
	consumer.Consumer()

}
