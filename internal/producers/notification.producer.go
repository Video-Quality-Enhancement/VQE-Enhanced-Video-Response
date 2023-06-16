package producers

import (
	"context"
	"encoding/json"
	"strings"
	"time"

	"github.com/Video-Quality-Enhancement/VQE-Enhanced-Video-Response/internal/config"
	"github.com/Video-Quality-Enhancement/VQE-Enhanced-Video-Response/internal/models"
	amqp "github.com/rabbitmq/amqp091-go"
	"golang.org/x/exp/slog"
)

type NotificationProducer interface {
	PublishNotification(request *models.EnhancedVideoNotifyRequest, notificationInterfacses []string) error
}

type notificationProducer struct {
	conn         config.AMQPconnection
	exchange     string
	exchangeType string
}

func NewNotificationProducer(conn config.AMQPconnection) NotificationProducer {

	exchange := "enhanced.video.notification"
	exchangeType := "topic"

	return &notificationProducer{
		conn:         conn,
		exchange:     exchange,
		exchangeType: exchangeType,
	}
}

func (producer *notificationProducer) PublishNotification(request *models.EnhancedVideoNotifyRequest, notificationInterfacses []string) error {

	ch, err := producer.conn.NewChannel()
	if err != nil {
		slog.Error("Failed to open a channel", "error", err)
		return err
	}
	defer ch.Close()

	err = ch.ExchangeDeclare(
		producer.exchange,
		producer.exchangeType,
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

	key := strings.Join(notificationInterfacses, ".")
	if key == "" {
		slog.Debug("No notification interfaces found, not publishing notification", "requestId", request.RequestId, "userId", request.UserId)
		return nil
	}

	err = ch.PublishWithContext(
		ctx,
		producer.exchange,
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
