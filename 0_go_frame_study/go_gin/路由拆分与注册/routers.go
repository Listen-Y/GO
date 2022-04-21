package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func helloHandler2(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "Hello www.topgoer.com!",
	})
}

func setupRouter() *gin.Engine {
	r := gin.Default()
	r.GET("/topgoer", helloHandler2)
	return r
}
