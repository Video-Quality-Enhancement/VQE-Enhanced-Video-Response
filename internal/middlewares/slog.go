package middlewares

import (
	"context"
	"time"

	"github.com/Video-Quality-Enhancement/VQE-Response-Producer/internal/config"
	"github.com/Video-Quality-Enhancement/VQE-Response-Producer/internal/utils"
	"github.com/Video-Quality-Enhancement/VQE-Response-Producer/internal/utils/seq"
	"golang.org/x/exp/slog"
)

func JSONlogger() seq.HandlerFunc {
	return func(c *seq.Context) {
		start := time.Now()

		c.Next()

		end := time.Now()
		latency := end.Sub(start)

		// make changes with the attributes for taking the producer and consumer in consern
		attributes := []slog.Attr{
			slog.String("gin-env", config.GetEnv("GIN_ENV", "development")),
			slog.String("service-name", config.GetEnv("SERVICE_NAME", "vqe-response-producer")),
			slog.String("user-id", utils.GetUserId(c)),
			slog.Int("status", c.Status()),
			slog.Duration("latency", latency),
			slog.String("request-id", utils.GetRequestID(c)),
		}

		switch {
		case c.Status() == 500:
			slog.LogAttrs(context.Background(), slog.LevelError, c.ErrorMessage(), attributes...)
		default:
			slog.LogAttrs(context.Background(), slog.LevelInfo, "Incoming request", attributes...)
		}

	}
}
