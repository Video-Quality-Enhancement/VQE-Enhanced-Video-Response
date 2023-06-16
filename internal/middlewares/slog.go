package middlewares

import (
	"context"
	"os"
	"time"

	"github.com/Video-Quality-Enhancement/VQE-Enhanced-Video-Response/internal/utils"
	"github.com/Video-Quality-Enhancement/VQE-Enhanced-Video-Response/internal/utils/tasks"
	"golang.org/x/exp/slog"
)

func JSONlogger() tasks.HandlerFunc {
	return func(c *tasks.Context) {
		start := time.Now()

		c.Next()

		end := time.Now()
		latency := end.Sub(start)

		userId, _ := utils.GetUserId(c)
		requestId, _ := utils.GetRequestID(c)

		// make changes with the attributes for taking the producer and consumer in consern
		attributes := []slog.Attr{
			slog.String("env", os.Getenv("ENV")),
			slog.String("service-name", os.Getenv("SERVICE_NAME")),
			slog.String("user-id", userId),
			slog.Int("status", c.Status()),
			slog.Duration("latency", latency),
			slog.String("request-id", requestId),
		}

		switch {
		case c.Status() == 500:
			slog.LogAttrs(context.Background(), slog.LevelError, c.ErrorMessage(), attributes...)
		default:
			slog.LogAttrs(context.Background(), slog.LevelInfo, "Incoming request", attributes...)
		}

	}
}
