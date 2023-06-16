package middlewares

import (
	"encoding/json"

	"github.com/Video-Quality-Enhancement/VQE-Enhanced-Video-Response/internal/models"
	"github.com/Video-Quality-Enhancement/VQE-Enhanced-Video-Response/internal/utils/tasks"
)

func SetEnhancedVideoProperties() tasks.HandlerFunc {
	return func(c *tasks.Context) {

		var response models.EnhancedVideoResponse
		err := json.Unmarshal(c.Delivery.Body, &response)
		if err != nil {
			c.Failure(err)
			return
		}
		c.Set("x-enhanced-video-response", &response)
		c.Set("X-Request-ID", response.RequestId)
		c.Set("X-User-ID", response.UserId)
		c.Next()

	}
}
