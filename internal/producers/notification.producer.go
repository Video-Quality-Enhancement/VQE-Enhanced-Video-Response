package producers

import (
	"context"
	"encoding/json"
	"strings"
	"time"

	"github.com/Video-Quality-Enhancement/VQE-Response-Producer/internal/models"
	amqp "github.com/rabbitmq/amqp091-go"
	"golang.org/x/exp/slog"
)

type NotificationProducer interface {
	PublishNotification(request *models.EnhancedVideoNotifyRequest) error
}

type notificationProducer struct {
	ch *amqp.Channel
}

func NewNotificationProducer(ch *amqp.Channel) NotificationProducer {
	return &notificationProducer{
		ch: ch,
	}
}

func (producer *notificationProducer) PublishNotification(request *models.EnhancedVideoNotifyRequest) error {

	exchange := "enhanced.video.notification"
	err := producer.ch.ExchangeDeclare(
		exchange,
		"topic",
		true,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		slog.Error("Failed to declare an exchange", "err", err)
		return err
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	body, err := json.Marshal(request)
	if err != nil {
		slog.Error("Failed to marshal video object", "err", err)
		return err
	}

	key := strings.Join(request.ResponseInterfaces, ".")
	if key == "" {
		slog.Debug("No response interfaces found, not publishing notification", "requestId", request.RequestId)
		return nil
	}

	err = producer.ch.PublishWithContext(
		ctx,
		exchange,
		key,
		false,
		false,
		amqp.Publishing{
			ContentType: "application/json",
			Body:        body,
		},
	)
	if err != nil {
		slog.Error("Failed to publish a message", "err", err)
		return err
	}

	slog.Debug("Message Published", "body", body)
	return nil

}
