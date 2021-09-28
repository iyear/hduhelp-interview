package middleware

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"os"
)

func Recovery() gin.HandlerFunc {
	return gin.CustomRecoveryWithWriter(os.Stderr, func(c *gin.Context, err interface{}) {
		zap.S().Errorw("recover", "error", err,
			"uri", c.Request.RequestURI)
	})
}
