package consumers

import (
	"os"
	"os/signal"
	"syscall"

	"github.com/Video-Quality-Enhancement/VQE-Response-Producer/internal/config"
	"github.com/Video-Quality-Enhancement/VQE-Response-Producer/internal/middlewares"
	"github.com/Video-Quality-Enhancement/VQE-Response-Producer/internal/utils/tasks"
	"github.com/Video-Quality-Enhancement/VQE-Response-Producer/internal/video/handlers"
	"github.com/Video-Quality-Enhancement/VQE-Response-Producer/internal/video/services"
	amqp "github.com/rabbitmq/amqp091-go"
	"golang.org/x/exp/slog"
)

type EnhancedVideoConsumer interface {
	Consumer() error
}

type enhancedVideoConsumer struct {
	conn    config.AMQPconnection
	service services.EnhancedVideoService
}

func NewEnhancedVideoConsumer(conn config.AMQPconnection, service services.EnhancedVideoService) EnhancedVideoConsumer {
	return &enhancedVideoConsumer{
		conn:    conn,
		service: service,
	}
}

func (consumer *enhancedVideoConsumer) Consumer() error {

	ch, err := consumer.conn.NewChannel()
	if err != nil {
		slog.Error("Failed to open a channel", "err", err)
		return err
	}
	defer ch.Close()

	q, err := ch.QueueDeclare(
		"enhanced.video", // name
		false,            // durable
		false,            // delete when unused
		false,            // exclusive
		false,            // no-wait
		nil,              // arguments
	)
	if err != nil {
		slog.Error("Failed to declare a queue", "err", err)
		return err
	}

	msgs, err := ch.Consume(
		q.Name, // queue
		"",     // consumer
		true,   // auto-ack
		false,  // exclusive
		false,  // no-local
		false,  // no-wait
		nil,    // args
	)
	if err != nil {
		slog.Error("Failed to declare a queue", "err", err)
		return err
	}

	var doneCh = make(chan struct{})

	go func(msgs <-chan amqp.Delivery, doneCh <-chan struct{}, service services.EnhancedVideoService) {
		for {
			select {
			case d := <-msgs:

				tasks.NewTask().Perform(
					middlewares.JSONlogger(),
					middlewares.SetEnhancedVideoProperties(d),
					handlers.EnhancedVideoHandler(service),
				)

				slog.Debug("Message Consumed", "body", d.Body)

			case <-doneCh:
				slog.Debug("Exiting the consumer goroutine")
				return
			}
		}
	}(msgs, doneCh, consumer.service)

	sigCh := make(chan os.Signal, 1)
	signal.Notify(sigCh, syscall.SIGINT, syscall.SIGTERM)

	slog.Debug("Waiting for messages. To exit press CTRL+C")

	<-sigCh

	doneCh <- struct{}{}

	slog.Debug("Exiting the consumer")

	return nil

}
