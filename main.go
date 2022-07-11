package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/llussy/go-utils/jwt"
	"github.com/llussy/go-utils/middleware"
)

func main() {

	user := jwt.User{
		ID:       1,
		Username: "llussy",
		Email:    "1299310393@qq.com",
		Password: "123456",
	}

	info, err := jwt.GenerateToken(user)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(info)

	a, err := jwt.ParseToken(info)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(a)

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
