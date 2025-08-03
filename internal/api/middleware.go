package api

import (
	"log"
	"net/http"

	"github.com/fareeza05/Code-Review-Assistant/internal/models"
	"github.com/gin-gonic/gin"
)

func ErrorHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Recover from any panic to prevent server crash
		defer func() {
			if err := recover(); err != nil {
				log.Printf("Unhandled error: %v", err)

				c.AbortWithStatusJSON(http.StatusInternalServerError, models.ErrorResponse{
					Error: "Internal server error",
				})
			}
		}()

		// Continue processing the request
		c.Next()

		// Check for manually set errors (optional)
		if len(c.Errors) > 0 {
			log.Printf("Handled error: %v", c.Errors.String())

			c.JSON(http.StatusBadRequest, models.ErrorResponse{
				Error: c.Errors.String(),
			})
		}
	}

}
