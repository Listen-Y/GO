package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

// 下面最基础的gin路由注册方式，适用于路由条目比较少的简单项目或者项目demo。
func helloHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "Hello www.topgoer.com!",
	})
}

func main1() {
	r := gin.Default()
	r.GET("/topgoer", helloHandler)
	if err := r.Run(":8000"); err != nil {
		fmt.Println("startup service failed, err:%v\n", err)
	}
}

// 当项目的规模增大后就不太适合继续在项目的main.go
//文件中去实现路由注册相关逻辑了，
//我们会倾向于把路由部分的代码都拆分出来，
//形成一个单独的文件或包：
// routers.go文件中定义并注册路由信息
// 此时main.go中调用上面定义好的setupRouter函数
func main() {
	r := setupRouter()
	if err := r.Run(":8000"); err != nil {
		fmt.Println("startup service failed, err:%v\n", err)
	}
}

// 当我们的业务规模继续膨胀，单独的一个routers文件或包已经满足不了我们的需求了
// routers/blog.go中添加一个LoadBlog的函数
// routers/shop.go中添加一个LoadShop的函数
// main方法会变成这样
/*func main() {
	r := gin.Default()
	routers.LoadBlog(r)
	routers.LoadShop(r)
	if err := r.Run(); err != nil {
		fmt.Println("startup service failed, err:%v\n", err)
	}
}*/

// 有时候项目规模实在太大，那么我们就更倾向于把业务拆分的更详细一些，例如把不同的业务代码拆分成不同的APP。
