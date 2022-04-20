package main

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"strings"
)

func main1() {
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

// 可以通过Context的Param方法来获取API参数
func main2() {
	r := gin.Default()
	r.GET("/user/:name/*action", func(c *gin.Context) {
		name := c.Param("name")
		action := c.Param("action")
		//截取/
		action = strings.Trim(action, "/")
		c.String(http.StatusOK, name+" is "+action)
	})
	//默认为监听8080端口
	r.Run(":8000")
}

// URL参数可以通过DefaultQuery()或Query()方法获取
//DefaultQuery()若参数不存在，返回默认值，Query()若不存在，返回空串
func main() {

}
