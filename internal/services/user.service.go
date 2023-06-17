package services

import (
	"github.com/Video-Quality-Enhancement/VQE-Enhanced-Video-Response/internal/config"
	"github.com/Video-Quality-Enhancement/VQE-Enhanced-Video-Response/internal/repositories"
	"golang.org/x/exp/slog"
)

type UserService interface {
	GetNotificationInterfaces(userId string) ([]string, error)
	GetEmail(userId string) (string, error)
	GetFCMTokens(userId string) ([]string, error)
}

type userService struct {
	repository     repositories.UserRepository
	firebaseClient config.FirebaseClient
}

func NewUserService(repository repositories.UserRepository, firebaseClient config.FirebaseClient) UserService {
	return &userService{repository, firebaseClient}
}

func (service *userService) GetNotificationInterfaces(userId string) ([]string, error) {

	notificationInterfaces, err := service.repository.FindNotificationInterfaces(userId)
	if err != nil {
		slog.Error("Error getting notification interfaces", "userId", userId)
		return nil, err
	}

	slog.Debug("Got notification interfaces", "userId", userId, "notificationInterfaces", notificationInterfaces)
	return notificationInterfaces, nil

}

func (service *userService) GetEmail(userId string) (string, error) {

	email, err := service.firebaseClient.GetEmail(userId)
	if err != nil {
		slog.Error("error getting user email", "error", err)
		return "", err
	}

	slog.Debug("got email of user", "userId", userId)
	return email, nil
}

func (service *userService) GetFCMTokens(userId string) ([]string, error) {

	fcmTokens, err := service.repository.FindFCMTokens(userId)
	if err != nil {
		slog.Error("Error getting FCM tokens", "userId", userId)
		return nil, err
	}

	slog.Debug("Got FCM tokens", "userId", userId, "fcmTokens", fcmTokens)
	return fcmTokens, nil

}
