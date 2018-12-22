package main

import "github.com/gin-gonic/gin"

var (
  Version  string
  Revision string
)

func main() {
	fmt.Print("Version:" + Version + " Revision:" + Reviison)

  // Global middleware
	r.Use(gin.Logger())

	// Routing
	router := gin.Default()
	router.GET("/api/greeting", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"message": "hello, world",
		})
	})
	r.Run()
}
