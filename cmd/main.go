package main

import (
	"github.com/Video-Quality-Enhancement/VQE-Enhanced-Video-Response/internal/app"
	"github.com/Video-Quality-Enhancement/VQE-Enhanced-Video-Response/internal/config"
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
	defer conn.DisconnectAll()

	firebaseClient := config.NewFirebaseClient()

	app.SetUpApp(db, conn, firebaseClient)

}
