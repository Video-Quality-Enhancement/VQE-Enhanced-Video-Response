package repositories

import (
	"context"
	"time"

	"github.com/Video-Quality-Enhancement/VQE-Enhanced-Video-Response/internal/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"golang.org/x/exp/slog"
)

type UserRepository interface {
	FindNotificationInterfaces(userId string) ([]string, error)
	FindFCMTokens(userId string) ([]string, error)
}

type userRepository struct {
	collection *mongo.Collection
}

func NewUserRepository(collection *mongo.Collection) UserRepository {
	return &userRepository{collection}
}

func (r *userRepository) FindNotificationInterfaces(userId string) ([]string, error) {

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	filter := bson.M{"userId": userId}
	opts := options.FindOne().SetProjection(bson.M{"notificationInterfaces": 1})

	var request models.NotificationInterfacesRequest
	err := r.collection.FindOne(ctx, filter, opts).Decode(&request)

	if err != nil {
		slog.Error("Failed to find Notification Interfaces", "error", err, "userId", userId)
		return nil, err
	}

	slog.Debug("Found Notification Interfaces", "userId", userId, "notificationInterfaces", request.NotificationInterfaces)
	return request.NotificationInterfaces, nil

}

func (r *userRepository) FindFCMTokens(userId string) ([]string, error) {

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	filter := bson.M{"userId": userId}
	opts := options.FindOne().SetProjection(bson.M{"fcmTokens": 1})

	var request models.FCMtokensRequest
	err := r.collection.FindOne(ctx, filter, opts).Decode(&request)

	if err != nil {
		slog.Error("Failed to find FCM Tokens", "error", err, "userId", userId)
		return nil, err
	}

	slog.Debug("Found FCM Tokens", "userId", userId, "fcmTokens", request.FCMtokens)
	return request.FCMtokens, nil

}
