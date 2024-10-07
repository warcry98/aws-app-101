package main

import (
	"backend/api/login"
	"backend/api/signup"
	"backend/api/validate"
	"backend/initializer"
	"backend/lib/middleware"

	"github.com/gin-gonic/gin"
)

func main() {
	dataEnv := initializer.LoadEnv()
	DB := initializer.ConnectToDB(dataEnv)
	initializer.SyncDatabase(DB)

	r := gin.New()
	r.Use(middleware.ApiMiddleware(DB))

	r.POST("/signup", signup.Controller)
	r.POST("/login", login.Controller)
	r.GET("/validate", middleware.RequireAuth, validate.Controller)

	r.Run()
}
