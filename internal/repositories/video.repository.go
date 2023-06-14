package repositories

import (
	"context"
	"time"

	"github.com/Video-Quality-Enhancement/VQE-Response-Producer/internal/models"
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
	FindByRequestId(userId, requestId string) (*models.EnhancedVideoNotifyRequest, error)
}

type enhancedVideoRepository struct {
	collection *mongo.Collection
}

func NewEnhancedVideoRepository(collection *mongo.Collection) EnhancedVideoRepository {
	return &enhancedVideoRepository{collection}
}

func (repository *enhancedVideoRepository) FindByRequestId(userId, requestId string) (*models.EnhancedVideoNotifyRequest, error) {

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	filter := bson.M{"userId": userId, "requestId": requestId}

	var request models.EnhancedVideoNotifyRequest
	err := repository.collection.FindOne(ctx, filter).Decode(&request)
	if err != nil {
		slog.Error("Error finding enhanced video", "requestId", requestId)
		return nil, err
	}

	slog.Debug("Found enhanced video", "requestId", request.RequestId)
	return &request, nil

}

func (repository *enhancedVideoRepository) Update(response *models.EnhancedVideoResponse) error {

	response.UpdatedAt = time.Now().UTC()

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	filter := bson.M{"userId": response.UserId, "requestId": response.RequestId}
	update := bson.D{
		{Key: "$set", Value: bson.M{
			"enhancedVideoUrl":     response.EnhancedVideoUrl,
			"enhancedVideoQuality": response.EnhancedVideoQuality,
			"status":               response.Status,
			"statusMessage":        response.StatusMessage,
			"updatedAt":            response.UpdatedAt,
		}}}

	updatedResult, err := repository.collection.UpdateOne(ctx, filter, update)

	if err != nil {
		slog.Error("Error updating video", "requestId", response.RequestId)
		return err
	}

	slog.Debug("Updated video", "requestId", response.RequestId, "updatedCount", updatedResult.ModifiedCount)
	return nil

}
