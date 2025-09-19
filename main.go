package main

import (
	"log"
	"main/Controller"
	"main/intializers"
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

	r.POST("/post", Controller.Postingfrom)
	r.POST("/postt", Controller.Postingto)
	r.POST("/transfer", Controller.TransferHandler)
	r.GET("/Get/:id", Controller.Get)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"

	}
	r.Run(":" + port)

}
