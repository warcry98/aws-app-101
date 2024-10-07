package main

import (
	"backend/api/login"
	"backend/api/signup"
	"backend/api/validate"
	"backend/initializer"
	"backend/lib/middleware"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	dataEnv := initializer.LoadEnv()
	DB := initializer.ConnectToDB(dataEnv)
	initializer.SyncDatabase(DB)

	r := gin.New()
	config := cors.DefaultConfig()
	config.AllowAllOrigins = true
	config.AllowHeaders = []string{"Origin", "Content-Length", "Content-Type", "Authorization"}
	config.AllowCredentials = true
	r.Use(cors.New(config))
	r.Use(middleware.ApiMiddleware(DB))

	api := r.Group("api")
	api.POST("/signup", signup.Controller)
	api.POST("/login", login.Controller)
	api.GET("/validate", middleware.RequireAuth, validate.Controller)

	r.Run()
}
