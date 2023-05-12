package middlewares

import "github.com/Video-Quality-Enhancement/VQE-Response-Producer/internal/utils/seq"

// instead of set requestId, make it set properties and then set requestId, userId and others
func SetRequestID() seq.HandlerFunc {
	return func(c *seq.Context) {

		requestID := "" // get this from the request
		c.Set("X-Request-ID", requestID)
		c.Next()

	}
}
