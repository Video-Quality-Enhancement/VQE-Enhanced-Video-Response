package producers

import (
	"context"
	"encoding/json"
	"strings"
	"time"

	"github.com/Video-Quality-Enhancement/VQE-Response-Producer/internal/config"
	"github.com/Video-Quality-Enhancement/VQE-Response-Producer/internal/video/models"
	amqp "github.com/rabbitmq/amqp091-go"
	"golang.org/x/exp/slog"
)

type NotificationProducer interface {
	PublishNotification(request *models.EnhancedVideoNotifyRequest) error
}

type notificationProducer struct {
	conn config.AMQPconnection
}

func NewNotificationProducer() NotificationProducer {
	return &notificationProducer{
		conn: config.NewAMQPconnection(),
	}
}

func (producer *notificationProducer) PublishNotification(request *models.EnhancedVideoNotifyRequest) error {

	ch, err := producer.conn.NewChannel()
	if err != nil {
		slog.Error("Failed to open a channel", "err", err)
		return err
	}
	defer ch.Close()

	err = ch.ExchangeDeclare(
		"notification",
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

	// ? Should I pass the correlation id also
	key := "notify." + strings.Join(request.ResponseInterfaces, ".")
	err = ch.PublishWithContext(
		ctx,
		"notification",
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
