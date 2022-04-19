package main

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

func main() {
	// 1.创建路由
	r := gin.Default()
	// 2.绑定路由规则，执行的函数
	// gin.Context，封装了request和response
	r.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "hello World!")
	})
	// 3.监听端口，默认在8080
	// Run("里面不指定端口号默认为8080")
	r.POST("/xxxpost", func(context *gin.Context) {
		context.String(http.StatusOK, "is a post")
	})
	r.PUT("/xxxput")
	err := r.Run(":8000")
	if err != nil {
		log.Fatalln("error")
	}
	log.Println("success")
}
