package main

//Casbin是用于Golang项目的功能强大且高效的开源访问控制库。

/*
Casbin的作用：

以经典{subject, object, action}形式或您定义的自定义形式实施策略，同时支持允许和拒绝授权。
处理访问控制模型及其策略的存储。
管理角色用户映射和角色角色映射（RBAC中的角色层次结构）。
支持内置的超级用户，例如root或administrator。超级用户可以在没有显式权限的情况下执行任何操作。
多个内置运算符支持规则匹配。例如，keyMatch可以将资源键映射/foo/bar到模式/foo*。
Casbin不执行的操作：

身份验证（又名验证username以及password用户登录时）
管理用户或角色列表。我相信项目本身管理这些实体会更方便。用户通常具有其密码，而Casbin并非设计为密码容器。但是，Casbin存储RBAC方案的用户角色映射。
*/

import (
	"fmt"

	"github.com/casbin/casbin"
	xormadapter "github.com/casbin/xorm-adapter"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	// 要使用自己定义的数据库rbac_db,最后的true很重要.默认为false,使用缺省的数据库名casbin,不存在则创建
	a := xormadapter.NewAdapter("mysql", "root:root@tcp(127.0.0.1:3306)/goblog?charset=utf8", true)

	// 可以在下面的文件中添加额外的用户信息
	enforcer := casbin.NewEnforcer("./rbac_models.conf", a)

	//从DB加载策略
	enforcer.LoadPolicy()

	//获取router路由对象
	r := gin.New()

	r.POST("/api/v1/add", func(c *gin.Context) {
		fmt.Println("增加Policy")
		if ok := enforcer.AddPolicy("admin", "/api/v1/hello", "GET"); !ok {
			fmt.Println("Policy已经存在")
		} else {
			fmt.Println("增加成功")
		}
	})
	//删除policy
	r.DELETE("/api/v1/delete", func(c *gin.Context) {
		fmt.Println("删除Policy")
		if ok := enforcer.RemovePolicy("admin", "/api/v1/hello", "GET"); !ok {
			fmt.Println("Policy不存在")
		} else {
			fmt.Println("删除成功")
		}
	})
	//获取policy
	r.GET("/api/v1/get", func(c *gin.Context) {
		fmt.Println("查看policy")
		list := enforcer.GetPolicy()
		for _, vlist := range list {
			for _, v := range vlist {
				fmt.Printf("value: %s, ", v)
			}
		}
	})
	//使用自定义拦截器中间件
	r.Use(Authorize(enforcer))
	//创建请求
	r.GET("/api/v1/hello", func(c *gin.Context) {
		fmt.Println("Hello 接收到GET请求..")
	})

	r.Run(":9000") //参数为空 默认监听8080端口
}

//拦截器
func Authorize(e *casbin.Enforcer) gin.HandlerFunc {

	return func(c *gin.Context) {

		//获取请求的URI
		obj := c.Request.URL.RequestURI()
		//获取请求方法
		act := c.Request.Method
		//获取用户的角色
		sub := "admin"

		//判断策略中是否存在
		if ok := e.Enforce(sub, obj, act); ok {
			fmt.Println("恭喜您,权限验证通过")
			c.Next()
		} else {
			fmt.Println("很遗憾,权限验证没有通过")
			c.Abort()
		}
	}
}
