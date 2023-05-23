package services

import (
	"github.com/Video-Quality-Enhancement/VQE-Response-Producer/internal/models"
	"github.com/Video-Quality-Enhancement/VQE-Response-Producer/internal/producers"
	"github.com/Video-Quality-Enhancement/VQE-Response-Producer/internal/repositories"
	amqp "github.com/rabbitmq/amqp091-go"
	"golang.org/x/exp/slog"
)

type EnhancedVideoService interface {
	OnVideoEnhancementComplete(response *models.EnhancedVideoResponse) error
}

type videoEnhanceService struct {
	repository           repositories.EnhancedVideoRepository
	notificationProducer producers.NotificationProducer
}

func NewEnhancedVideoService(repository repositories.EnhancedVideoRepository, ch *amqp.Channel) EnhancedVideoService {

	return &videoEnhanceService{
		repository:           repository,
		notificationProducer: producers.NewNotificationProducer(ch),
	}

}

func (service *videoEnhanceService) getNotifyRequest(userId, requestId string) (*models.EnhancedVideoNotifyRequest, error) {

	notifyRequest, err := service.repository.FindByRequestId(userId, requestId)
	if err != nil {
		slog.Error("Error getting notify request", "requestId", requestId)
		return nil, err
	}

	slog.Debug("Got notify request", "requestId", requestId)
	return notifyRequest, nil

}

func (service *videoEnhanceService) OnVideoEnhancementComplete(response *models.EnhancedVideoResponse) error {

	err := service.repository.Update(response)
	if err != nil {
		slog.Error("Error updating video", "requestId", response.RequestId)
		return err
	}

	notifyRequest, err := service.getNotifyRequest(response.UserId, response.RequestId)
	if err != nil {
		slog.Error("Error getting notify request", "requestId", response.RequestId)
		return err
	}

	err = service.notificationProducer.PublishNotification(notifyRequest) // not running this in a serparate goroutine coz i will run the enhanced video consumer in a separate goroutine which calls this method and even record the time taken to update and publish using the slog middleware
	if err != nil {
		slog.Error("Error publishing notification", "requestId", response.RequestId)
		return err
	}

	slog.Debug("Updated video", "requestId", response.RequestId)
	return nil

}
