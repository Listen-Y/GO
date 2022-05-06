package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"time"
)

// 所有请求都经过此中间件
// 定义中间
func MiddleWare1() gin.HandlerFunc {
	return func(c *gin.Context) {
		t := time.Now()
		fmt.Println("中间件开始执行了")
		// 设置变量到Context的key中，可以通过Get()取
		c.Set("request", "中间件")
		status := c.Writer.Status()
		fmt.Println("中间件执行完毕", status)
		t2 := time.Since(t)
		fmt.Println("time:", t2)
	}
}
func main1() {
	// 1.创建路由
	// 默认使用了2个中间件Logger(), Recovery()
	r := gin.Default()
	// 注册中间件
	r.Use(MiddleWare1())
	// {}为了代码规范
	{
		r.GET("/ce", func(c *gin.Context) {
			// 取值
			req, _ := c.Get("request")
			fmt.Println("request:", req)
			// 页面接收
			c.JSON(200, gin.H{"request": req})
		})

	}
	r.Run()
}

// Next方法 可以快速切换中间件函数的执行到业务函数的执行中，然后最后执行剩下的中间件函数代码
// 定义中间
func MiddleWare2() gin.HandlerFunc {
	return func(c *gin.Context) {
		t := time.Now()
		fmt.Println("中间件开始执行了")
		// 设置变量到Context的key中，可以通过Get()取
		c.Set("request", "中间件")
		// 执行函数
		c.Next()
		// 中间件执行完后续的一些事情
		status := c.Writer.Status()
		fmt.Println("中间件执行完毕", status)
		t2 := time.Since(t)
		fmt.Println("time:", t2)
	}
}

func main2() {
	// 1.创建路由
	// 默认使用了2个中间件Logger(), Recovery()
	r := gin.Default()
	// 注册中间件
	r.Use(MiddleWare2())
	// {}为了代码规范
	{
		r.GET("/ce", func(c *gin.Context) {
			// 取值
			req, _ := c.Get("request")
			fmt.Println("request:", req)
			// 页面接收
			c.JSON(200, gin.H{"request": req})
		})

	}
	r.Run()
}

// 局部中间件
// 定义中间
func MiddleWare3() gin.HandlerFunc {
	return func(c *gin.Context) {
		t := time.Now()
		fmt.Println("中间件开始执行了")
		// 设置变量到Context的key中，可以通过Get()取
		c.Set("request", "中间件")
		// 执行函数
		c.Next()
		// 中间件执行完后续的一些事情
		status := c.Writer.Status()
		fmt.Println("中间件执行完毕", status)
		// 统计用时
		t2 := time.Since(t)
		fmt.Println("time:", t2)
	}
}

func main3() {
	// 1.创建路由
	// 默认使用了2个中间件Logger(), Recovery()
	r := gin.Default()
	//局部中间键使用
	r.GET("/ce", MiddleWare3(), func(c *gin.Context) {
		// 取值
		req, _ := c.Get("request")
		fmt.Println("request:", req)
		// 页面接收
		c.JSON(200, gin.H{"request": req})
	})
	r.Run()
}

// 中间件推荐
/*
列表
谷歌翻译欢迎查看原文 https://github.com/gin-gonic/contrib/blob/master/README.md

RestGate - REST API端点的安全身份验证
staticbin - 用于从二进制数据提供静态文件的中间件/处理程序
gin-cors - CORS杜松子酒的官方中间件
gin-csrf - CSRF保护
gin-health - 通过gocraft/health报告的中间件
gin-merry - 带有上下文的漂亮 打印 错误的中间件
gin-revision - 用于Gin框架的修订中间件
gin-jwt - 用于Gin框架的JWT中间件
gin-sessions - 基于mongodb和mysql的会话中间件
gin-location - 用于公开服务器的主机名和方案的中间件
gin-nice-recovery - 紧急恢复中间件，可让您构建更好的用户体验
gin-limit - 限制同时请求；可以帮助增加交通流量
gin-limit-by-key - 一种内存中的中间件，用于通过自定义键和速率限制访问速率。
ez-gin-template - gin简单模板包装
gin-hydra - gin中间件Hydra
gin-glog - 旨在替代Gin的默认日志
gin-gomonitor - 用于通过Go-Monitor公开指标
gin-oauth2 - 用于OAuth2
static gin框架的替代静态资产处理程序。
xss-mw - XssMw是一种中间件，旨在从用户提交的输入中“自动删除XSS”
gin-helmet - 简单的安全中间件集合。
gin-jwt-session - 提供JWT / Session / Flash的中间件，易于使用，同时还提供必要的调整选项。也提供样品。
gin-template - 用于gin框架的html / template易于使用。
gin-redis-ip-limiter - 基于IP地址的请求限制器。它可以与redis和滑动窗口机制一起使用。
gin-method-override - _method受Ruby的同名机架启发而被POST形式参数覆盖的方法
gin-access-limit - limit-通过指定允许的源CIDR表示法的访问控制中间件。
gin-session - 用于Gin的Session中间件
gin-stats - 轻量级和有用的请求指标中间件
gin-statsd - 向statsd守护进程报告的Gin中间件
gin-health-check - check-用于Gin的健康检查中间件
gin-session-middleware - 一个有效，安全且易于使用的Go Session库。
ginception - 漂亮的例外页面
gin-inspector - 用于调查http请求的Gin中间件。
gin-dump - Gin中间件/处理程序，用于转储请求和响应的标头/正文。对调试应用程序非常有帮助。
go-gin-prometheus - Gin Prometheus metrics exporter
ginprom - Gin的Prometheus指标导出器
gin-go-metrics - Gin middleware to gather and store metrics using rcrowley/go-metrics
ginrpc - Gin 中间件/处理器自动绑定工具。通过像beego这样的注释路线来支持对象注册
*/
