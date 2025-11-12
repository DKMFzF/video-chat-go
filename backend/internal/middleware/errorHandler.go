package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type APIError struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

func ErrorHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next()
		err := c.Errors
		if len(err) > 0 {
			lastErr := err.Last()
			c.JSON(http.StatusInternalServerError, APIError{
				Code:    http.StatusInternalServerError,
				Message: lastErr.Error(),
			})

			c.Abort()
		}
	}
}
