package consumers

import (
	"os"
	"os/signal"
	"syscall"

	"github.com/Video-Quality-Enhancement/VQE-Response-Producer/internal/config"
	"github.com/Video-Quality-Enhancement/VQE-Response-Producer/internal/handlers"
	"github.com/Video-Quality-Enhancement/VQE-Response-Producer/internal/middlewares"
	"github.com/Video-Quality-Enhancement/VQE-Response-Producer/internal/services"
	"github.com/Video-Quality-Enhancement/VQE-Response-Producer/internal/utils/tasks"
	amqp "github.com/rabbitmq/amqp091-go"
	"golang.org/x/exp/slog"
)

type EnhancedVideoConsumer interface {
	Consume() error
}

type enhancedVideoConsumer struct {
	conn      config.AMQPconnection
	service   services.EnhancedVideoService
	queueName string
}

func NewEnhancedVideoConsumer(conn config.AMQPconnection, service services.EnhancedVideoService) EnhancedVideoConsumer {
	queueName := "enhanced.video"
	return &enhancedVideoConsumer{
		conn:      conn,
		service:   service,
		queueName: queueName,
	}
}

func (consumer *enhancedVideoConsumer) Consume() error {

	ch, err := consumer.conn.NewChannel()
	if err != nil {
		slog.Error("Failed to open a channel", "error", err)
		return err
	}
	defer ch.Close()

	q, err := ch.QueueDeclare(
		consumer.queueName, // name
		true,               // durable
		false,              // delete when unused
		false,              // exclusive
		false,              // no-wait
		nil,                // arguments
	)
	if err != nil {
		slog.Error("Failed to declare a queue", "err", err)
		return err
	}

	msgs, err := ch.Consume(
		q.Name, // queue
		"",     // consumer
		false,  // auto-ack
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

		task := tasks.NewTask()
		task.Activities(
			middlewares.JSONlogger(),
			middlewares.SetEnhancedVideoProperties(),
			handlers.EnhancedVideoHandler(service),
		)

		for {
			select {
			case msg := <-msgs:

				task.Perform(msg)
				slog.Debug("Message Consumed, msg is either acked or nacked")

			case <-doneCh:
				slog.Info("Exiting the consumer goroutine")
				return
			}
		}

	}(msgs, doneCh, consumer.service)

	sigCh := make(chan os.Signal, 1)
	signal.Notify(sigCh, syscall.SIGINT, syscall.SIGTERM)

	slog.Info("Waiting for messages. To exit press CTRL+C")

	<-sigCh

	doneCh <- struct{}{}

	slog.Info("Exiting the consumer")

	return nil

}
