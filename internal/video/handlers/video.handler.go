package handlers

import (
	"github.com/Video-Quality-Enhancement/VQE-Response-Producer/internal/utils"
	"github.com/Video-Quality-Enhancement/VQE-Response-Producer/internal/utils/tasks"
	"github.com/Video-Quality-Enhancement/VQE-Response-Producer/internal/video/services"
	"golang.org/x/exp/slog"
)

func EnhancedVideoHandler(service services.EnhancedVideoService) tasks.HandlerFunc {
	return func(c *tasks.Context) {

		response := utils.GetEnhancedVideoResponse(c)
		err := service.OnVideoEnhancementComplete(response)
		if err != nil {
			slog.Error("Error handling enhanced video", "requestId", response.RequestId, "err", err)
			c.Failure(err)
			return
		}

		slog.Debug("Enhanced video handled", "requestId", response.RequestId)
		c.Success()

	}
}
