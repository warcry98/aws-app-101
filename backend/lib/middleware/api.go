package middleware

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func ApiMiddleware(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Set("DB", db)
		c.Next()
	}
}
