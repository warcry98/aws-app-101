package signup

import (
	"backend/models"
	"database/sql"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

func Controller(c *gin.Context) {
	var body struct {
		Email    string `binding:"required,email"`
		Username string `binding:"required"`
		Fullname string
		Password string `binding:"required"`
	}

	if c.Bind(&body) != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to read body",
		})

		return
	}

	hash, err := bcrypt.GenerateFromPassword([]byte(body.Password), 10)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to hash password.",
		})
		return
	}

	var fullname sql.NullString
	if len(body.Fullname) == 0 {
		fullname = sql.NullString{String: "", Valid: false}
	} else {
		fullname = sql.NullString{String: body.Fullname, Valid: true}
	}

	user := models.User{
		Email:    body.Email,
		Username: body.Username,
		Password: string(hash),
		Fullname: fullname,
	}

	DB, ok := c.MustGet("DB").(*gorm.DB)
	if !ok {
		log.Fatalf("Failed to connect to DB.")
	}

	result := DB.Create(&user)

	if result.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to create user.",
		})
	}

	c.JSON(http.StatusOK, gin.H{})
}
