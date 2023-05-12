package repositories

import (
	"context"
	"time"

	"github.com/Video-Quality-Enhancement/VQE-Response-Producer/internal/video/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"golang.org/x/exp/slog"
)

type EnhancedVideoRepository interface {
	EnhancedVideoResponseRepository
	EnhancedVideoNotifyRequestRepository
}

type EnhancedVideoResponseRepository interface {
	Update(response *models.EnhancedVideoResponse) error
}

type EnhancedVideoNotifyRequestRepository interface {
	FindByRequestId(requestId string) (*models.EnhancedVideoNotifyRequest, error)
}

type enhancedVideoRepository struct {
	collection *mongo.Collection
}

func NewEnhancedVideoRepository(collection *mongo.Collection) EnhancedVideoRepository {
	return &enhancedVideoRepository{collection}
}

func (repository *enhancedVideoRepository) FindByRequestId(requestId string) (*models.EnhancedVideoNotifyRequest, error) {

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	var request models.EnhancedVideoNotifyRequest
	err := repository.collection.FindOne(ctx, models.EnhancedVideoNotifyRequest{RequestId: requestId}).Decode(&request)
	if err != nil {
		slog.Error("Error finding enhanced video", "requestId", requestId)
		return nil, err
	}

	slog.Debug("Found enhanced video", "requestId", request.RequestId)
	return &request, nil

}

func (repository *enhancedVideoRepository) Update(response *models.EnhancedVideoResponse) error {

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	_, err := repository.collection.UpdateOne(
		ctx,
		models.EnhancedVideoResponse{RequestId: response.RequestId},
		bson.D{{Key: "$set", Value: models.EnhancedVideoResponse{EnhancedVideoUri: response.EnhancedVideoUri}}},
	)

	if err != nil {
		slog.Error("Error updating video", "requestId", response.RequestId)
		return err
	}

	slog.Debug("Updated video", "requestId", response.RequestId)
	return nil

}
