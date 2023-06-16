package handlers

import (
	"github.com/Video-Quality-Enhancement/VQE-Enhanced-Video-Response/internal/services"
	"github.com/Video-Quality-Enhancement/VQE-Enhanced-Video-Response/internal/utils"
	"github.com/Video-Quality-Enhancement/VQE-Enhanced-Video-Response/internal/utils/tasks"
	"golang.org/x/exp/slog"
)

func EnhancedVideoHandler(service services.EnhancedVideoService) tasks.HandlerFunc {
	return func(c *tasks.Context) {

		response, err := utils.GetEnhancedVideoResponse(c)
		if err != nil {
			slog.Error("Error getting enhanced video response", "err", err)
			c.Failure(err)
			return
		}

		err = service.OnVideoEnhancementComplete(response)
		if err != nil {
			slog.Error("Error handling enhanced video", "requestId", response.RequestId, "err", err)
			c.Failure(err)
			return
		}

		slog.Debug("Enhanced video handled", "requestId", response.RequestId)
		c.Success()

	}
}
