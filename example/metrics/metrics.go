package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/llussy/go-utils/middleware"
)

func main() {

	g := gin.New()
	g.Use(middleware.Metrics(), gin.Recovery())

	g.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "ok",
		})
	})

	g.GET("/metrics", middleware.Monitor)

	g.Run(fmt.Sprintf(":%d", 8000))

}
