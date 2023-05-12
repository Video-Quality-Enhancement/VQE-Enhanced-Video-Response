package utils

import (
	"github.com/Video-Quality-Enhancement/VQE-Response-Producer/internal/utils/seq"
	"github.com/Video-Quality-Enhancement/VQE-Response-Producer/internal/video/models"
	"golang.org/x/exp/slog"
)

func SetEnhancedVideoResponse(c *seq.Context, response *models.EnhancedVideoResponse) {

	if response == nil {
		slog.Warn("Enhanced Video Response missing, cannot set enhanced video response")
	} else {
		c.Set("x-enhanced-video-response", response)
	}

}

func GetEnhancedVideoResponse(c *seq.Context) *models.EnhancedVideoResponse {

	response := c.Get("x-enhanced-video-response").(*models.EnhancedVideoResponse)

	if response == nil {
		slog.Error("Enhanced Video Response missing, cannot get enhanced video response")
		return nil
	}

	return response

}
