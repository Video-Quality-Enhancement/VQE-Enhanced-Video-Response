package utils

import (
	"github.com/Video-Quality-Enhancement/VQE-Response-Producer/internal/utils/seq"
	"golang.org/x/exp/slog"
)

func SetUserId(c *seq.Context, userId string) {

	if userId == "" {
		slog.Warn("User ID missing, cannot set userId")
	} else {
		c.Set("x-userId", userId)
	}

}

func GetUserId(c *seq.Context) string {

	userId := c.Get("x-userId").(string)

	if userId == "" {
		slog.Error("User ID missing, cannot get userId")
		return ""
	}

	return userId
}
