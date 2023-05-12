package handlers

import (
	"github.com/Video-Quality-Enhancement/VQE-Response-Producer/internal/utils/seq"
	"github.com/Video-Quality-Enhancement/VQE-Response-Producer/internal/video/services"
	"github.com/Video-Quality-Enhancement/VQE-Response-Producer/internal/video/utils"
	"golang.org/x/exp/slog"
)

func EnhancedVideoHandler(service services.EnhancedVideoService) seq.HandlerFunc {
	return func(c *seq.Context) {

		response := utils.GetEnhancedVideoResponse(c)
		err := service.OnVideoEnhancementComplete(response)
		if err != nil {
			slog.Error("Error handling enhanced video", "requestId", response.RequestId)
			c.Error(err)
			return
		}

		slog.Debug("Enhanced video handled", "requestId", response.RequestId)
		c.Success()

	}
}
