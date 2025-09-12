package main

import (
	"go-docker-demo/routes"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	routes.RegisterJokeRoutes(router)

	router.GET("/html", func(c *gin.Context) {
		c.Data(http.StatusOK, "text/html; charset=utf-8", []byte(`<h1>Hello, Gin! <span style="color:#55D7E5">ʕ◔ϖ◔ʔ</span></h1>`))
	})
	router.GET("/kiba", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "Hello, Kiba!",
		})
	})
	router.POST("/ping", func(c *gin.Context) {
		var req struct {
			Message string `json:"message"`
		}
		if err := c.BindJSON(&req); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON"})
			return
		}
		c.JSON(http.StatusOK, gin.H{
			"message": "Pong " + req.Message,
		})
	})
	router.Run(":4269")
}
