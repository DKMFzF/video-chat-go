package middleware

import (
	"strconv"
	"video-chat/internal/logger"

	"github.com/gin-gonic/gin"
)

func LoggerHandler(log *logger.ZapLogger) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next()

		log.Infof(
			"%s", "Request "+
				c.Request.RequestURI+
				" Response Code"+strconv.Itoa(c.Writer.Status()),
		)
	}
}
