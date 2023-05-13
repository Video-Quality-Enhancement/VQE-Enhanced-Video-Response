package middlewares

import (
	"encoding/json"

	"github.com/Video-Quality-Enhancement/VQE-Response-Producer/internal/utils/tasks"
	"github.com/Video-Quality-Enhancement/VQE-Response-Producer/internal/video/models"
	amqp "github.com/rabbitmq/amqp091-go"
)

// instead of set requestId, make it set properties and then set requestId, userId and others
func SetEnhancedVideoProperties(d amqp.Delivery) tasks.HandlerFunc {
	return func(c *tasks.Context) {

		var response models.EnhancedVideoResponse
		err := json.Unmarshal(d.Body, &response)
		if err != nil {
			c.AbortWithError(err)
			return
		}
		c.Set("x-enhanced-video-response", response)
		c.Set("X-Request-ID", response.RequestId)
		c.Set("X-User-ID", response.UserId)
		c.Next()

	}
}
