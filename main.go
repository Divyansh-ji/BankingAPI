package main

import (
	"log"
	"main/Controller"
	"main/intializers"
	"main/middleware"
	"os"

	"github.com/gin-gonic/gin"
)

func init() {
	intializers.LoadEnvVariables()
	intializers.ConnectToDB()

	intializers.SyncDatabase() // <--- add this

}
func main() {
	log.Println("app started")
	gin.SetMode(gin.DebugMode)
	r := gin.Default()

	authGroup := r.Group("/api")
	authGroup.Use(middleware.RequireAuth)
	{
		r.POST("/post", Controller.Postingfrom)
		r.POST("/postt", Controller.Postingto)
		r.POST("/transfer", Controller.TransferHandler)
		r.GET("/get/:id", Controller.Get)
	}
	r.POST("signup", Controller.SignUp)
	r.POST("login", Controller.Login)
	r.POST("/refreshToken", Controller.RefreshToken)
	r.GET("/logout", Controller.Logout)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"

	}
	r.Run(":" + port)

}
