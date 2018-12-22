package main

import (
	"fmt"
)

import (
	"github.com/fvbock/endless"
	"github.com/gin-gonic/gin"
)

var (
	Version  string
	Revision string
)

func setupRouter() *gin.Engine {
	router := gin.Default()

	// Global middleware
	router.Use(gin.Logger())

	// Routing
	router.StaticFile("/", "./index.html")

	router.GET("/api/greeting", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"message": "hello, world",
		})
	})

	return router
}

func main() {
	fmt.Println("Greetings Server : Version:" + Version + " Revision:" + Revision)

	endless.ListenAndServe(":8080", setupRouter())
}
