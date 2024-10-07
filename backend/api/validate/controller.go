package validate

import (
	"backend/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Controller(c *gin.Context) {
	user, ok := c.Get("user")
	if ok {
		data_user := user.(models.User)

		c.JSON(http.StatusOK, gin.H{
			"message": data_user.Username,
		})
	} else {
		c.AbortWithStatus(http.StatusUnauthorized)
	}
}
