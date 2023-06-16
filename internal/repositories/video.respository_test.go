package repositories_test

// import (
// 	"os"
// 	"testing"

// 	"github.com/Video-Quality-Enhancement/VQE-Response-Producer/internal/config"
// 	"github.com/Video-Quality-Enhancement/VQE-Response-Producer/internal/models"
// 	"github.com/Video-Quality-Enhancement/VQE-Response-Producer/internal/repositories"
// )

// func TestMain(m *testing.M) {

// 	os.Exit(m.Run())
// }

// func TestEnhancedVideoRepository(t *testing.T) {

// 	t.Setenv("MONGO_URI", "mongodb://localhost:27017")
// 	t.Setenv("MONGO_DB", "VQE")

// 	client := config.NewMongoClient()
// 	db := client.ConnectToDB()
// 	collection := db.Collection("videos")

// 	repository := repositories.NewEnhancedVideoRepository(collection)

// 	t.Run("Test Enhanced Video Response Repository", GetUpdateTest(repository))
// 	t.Run("Test Enhanced Video Notify Request Repository", GetFindByRequestIdTest(repository))

// 	t.Cleanup(client.Disconnect)

// }

// func GetFindByRequestIdTest(repository repositories.EnhancedVideoNotifyRequestRepository) func(t *testing.T) {

// 	return func(t *testing.T) {

// 		notifyRequest, err := repository.FindByRequestId("1234", "431793d2-d36d-4fcc-8598-85875fef3c2c")

// 		if err != nil {
// 			t.Error("Error finding Enhanced Video Notify Request", "err", err)
// 		} else {
// 			t.Log(notifyRequest)
// 		}

// 	}

// }

// func GetUpdateTest(repository repositories.EnhancedVideoResponseRepository) func(t *testing.T) {

// 	return func(t *testing.T) {

// 		var response = models.EnhancedVideoResponse{
// 			UserId:           "1234",
// 			RequestId:        "431793d2-d36d-4fcc-8598-85875fef3c2c",
// 			EnhancedVideoUrl: "https://download.samplelib.com/mp4/sample-5s.mp4",
// 			Status:           "COMPLETED",
// 			StatusMessage:    "Enhanced Video has been created",
// 		}

// 		err := repository.Update(&response)

// 		if err != nil {
// 			t.Error("Error updating Enhanced Video Response", "err", err)
// 		}

// 	}

// }
