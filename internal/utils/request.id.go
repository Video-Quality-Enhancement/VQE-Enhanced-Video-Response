package utils

import (
	"github.com/Video-Quality-Enhancement/VQE-Response-Producer/internal/utils/seq"
	"golang.org/x/exp/slog"
)

func GetRequestID(c *seq.Context) string {
	requestID := c.Get("X-Request-ID").(string)

	if requestID == "" {
		slog.Error("Request ID not found in header")
		return ""
	}

	return requestID
}
