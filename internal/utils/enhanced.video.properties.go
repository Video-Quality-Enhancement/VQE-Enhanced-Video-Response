package utils

import (
	"errors"

	"github.com/Video-Quality-Enhancement/VQE-Response-Producer/internal/models"
	"github.com/Video-Quality-Enhancement/VQE-Response-Producer/internal/utils/tasks"
	"golang.org/x/exp/slog"
)

func GetRequestID(c *tasks.Context) (string, error) {
	requestID := c.Get("X-Request-ID").(string)

	if requestID == "" {
		slog.Error("Request ID not found in header")
		return "", errors.New("request ID not found in header")
	}

	return requestID, nil
}

func GetUserId(c *tasks.Context) (string, error) {
	userId := c.Get("X-User-ID").(string)

	if userId == "" {
		slog.Error("User ID missing, cannot get userId")
		return "", errors.New("user ID missing, cannot get userId")
	}

	return userId, nil
}

func GetEnhancedVideoResponse(c *tasks.Context) (*models.EnhancedVideoResponse, error) {

	response := c.Get("x-enhanced-video-response").(*models.EnhancedVideoResponse)

	if response == nil {
		slog.Error("Enhanced Video Response missing, cannot get Enhanced Video Response")
		return nil, errors.New("enhanced Video Response missing, cannot get Enhanced Video Response")
	}

	return response, nil
}
