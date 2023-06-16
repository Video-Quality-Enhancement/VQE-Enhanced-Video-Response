package services

import (
	"github.com/Video-Quality-Enhancement/VQE-Enhanced-Video-Response/internal/models"
	"github.com/Video-Quality-Enhancement/VQE-Enhanced-Video-Response/internal/producers"
	"github.com/Video-Quality-Enhancement/VQE-Enhanced-Video-Response/internal/repositories"
	"github.com/Video-Quality-Enhancement/VQE-Enhanced-Video-Response/internal/services/gapi"
	"golang.org/x/exp/slog"
)

type EnhancedVideoService interface {
	OnVideoEnhancementComplete(response *models.EnhancedVideoResponse) error
}

type enhancedVideoService struct {
	repository           repositories.EnhancedVideoRepository
	userService          UserService
	storageService       gapi.GoogleCloudStorage
	notificationProducer producers.NotificationProducer
}

func NewEnhancedVideoService(
	repository repositories.EnhancedVideoRepository,
	userService UserService,
	producer producers.NotificationProducer,
) EnhancedVideoService {

	storageService := gapi.NewGoogleCloudStorage()

	return &enhancedVideoService{
		repository:           repository,
		userService:          userService,
		storageService:       storageService,
		notificationProducer: producer,
	}

}

func (service *enhancedVideoService) OnVideoEnhancementComplete(response *models.EnhancedVideoResponse) error {

	err := service.repository.Update(response)
	if err != nil {
		slog.Error("Error updating video", "requestId", response.RequestId)
		return err
	}

	email, err := service.userService.GetEmail(response.UserId)
	if err != nil {
		slog.Error("error getting email from user service", "requestId", response.RequestId, "useId", response.UserId)
		return err
	}

	filepath := response.RequestId + ".mp4"
	err = service.storageService.GrantAccess(filepath, email)
	if err != nil {
		slog.Error("error granting access to email", "filepath", filepath, "requestId", response.RequestId, "useId", response.UserId)
		return err
	}

	notificationInterfaces, err := service.userService.GetNotificationInterfaces(response.UserId)
	if err != nil {
		slog.Error("error getting notification interfaces from user service", "error", err, "requestId", response.RequestId, "useId", response.UserId)
		return err
	}

	notifyRequest := &models.EnhancedVideoNotifyRequest{
		RequestId:            response.RequestId,
		UserId:               response.UserId,
		EnhancedVideoUrl:     response.EnhancedVideoUrl,
		EnhancedVideoQuality: response.EnhancedVideoQuality,
		Status:               response.Status,
	}
	err = service.notificationProducer.PublishNotification(notifyRequest, notificationInterfaces) // not running this in a serparate goroutine coz i will run the enhanced video consumer in a separate goroutine which calls this method and even record the time taken to update and publish using the slog middleware
	if err != nil {
		slog.Error("Error publishing notification", "requestId", response.RequestId)
		return err
	}

	slog.Debug("Updated video", "requestId", response.RequestId)
	return nil

}
