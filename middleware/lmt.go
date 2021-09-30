package middleware

import (
	"github.com/didip/tollbooth"
	"github.com/didip/tollbooth/limiter"
	"github.com/gin-gonic/gin"
	"github.com/iyear/hduhelp-interview/api"
	"github.com/iyear/hduhelp-interview/conf"
	"github.com/iyear/hduhelp-interview/conf/e"
)

func RateLimit(lmt *limiter.Limiter) gin.HandlerFunc {
	return func(c *gin.Context) {
		err := tollbooth.LimitByRequest(lmt, c.Writer, c.Request)
		if err != nil {
			c.Abort()
			api.RespFmt(c, err.StatusCode, &api.Resp{
				Error:    e.RATE_LIMIT,
				Msg:      e.GetMsg(e.RATE_LIMIT),
				Redirect: "",
				Data:     nil,
			})
			return
		}
		c.Next()
	}
}
func RateLimitInit() *limiter.Limiter {
	return tollbooth.NewLimiter(float64(conf.App.Limit), nil)
}
