package utils

import (
	"github.com/Video-Quality-Enhancement/VQE-Response-Producer/internal/utils/tasks"
	"github.com/Video-Quality-Enhancement/VQE-Response-Producer/internal/video/models"
	"golang.org/x/exp/slog"
)

func GetRequestID(c *tasks.Context) string {
	requestID := c.Get("X-Request-ID").(string)

	if requestID == "" {
		slog.Error("Request ID not found in header")
		return ""
	}

	return requestID
}

func GetUserId(c *tasks.Context) string {

	userId := c.Get("X-User-ID").(string)

	if userId == "" {
		slog.Error("User ID missing, cannot get userId")
		return ""
	}

	return userId
}

func GetEnhancedVideoResponse(c *tasks.Context) *models.EnhancedVideoResponse {

	response := c.Get("x-enhanced-video-response").(*models.EnhancedVideoResponse)

	if response == nil {
		slog.Error("Enhanced Video Response missing, cannot get Enhanced Video Response")
		return nil
	}

	return response
}
