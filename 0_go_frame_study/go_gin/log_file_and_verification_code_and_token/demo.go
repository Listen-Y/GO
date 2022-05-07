package main

import (
	"bytes"
	"fmt"
	"github.com/dchest/captcha"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
	"io"
	"net/http"
	"os"
	"time"
)

// 日志文件
func main1() {
	gin.DisableConsoleColor()

	// Logging to a file.
	f, _ := os.Create("gin.log")
	gin.DefaultWriter = io.MultiWriter(f)

	// 如果需要同时将日志写入文件和控制台，请使用以下代码。
	// gin.DefaultWriter = io.MultiWriter(f, os.Stdout)
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.String(200, "pong")
	})
	r.Run()
}

// 验证码
/*
web端是怎么实现验证码的功能呢？

提供一个路由，先在session里写入键值对（k->v），把值写在图片上，然后生成图片，显示在浏览器上面
前端将图片中的内容发送给后后端，后端根据session中的k取得v，比对校验。如果通过继续下一步的逻辑，失败给出错误提示
*/

// 中间件，处理session
func Session(keyPairs string) gin.HandlerFunc {
	store := SessionConfig()
	return sessions.Sessions(keyPairs, store)
}
func SessionConfig() sessions.Store {
	sessionMaxAge := 3600
	sessionSecret := "topgoer"
	var store sessions.Store
	store = cookie.NewStore([]byte(sessionSecret))
	store.Options(sessions.Options{
		MaxAge: sessionMaxAge, //seconds
		Path:   "/",
	})
	return store
}

func Captcha(c *gin.Context, length ...int) {
	l := captcha.DefaultLen
	w, h := 107, 36
	if len(length) == 1 {
		l = length[0]
	}
	if len(length) == 2 {
		w = length[1]
	}
	if len(length) == 3 {
		h = length[2]
	}
	captchaId := captcha.NewLen(l)
	session := sessions.Default(c)
	session.Set("captcha", captchaId)
	_ = session.Save()
	_ = Serve(c.Writer, c.Request, captchaId, ".png", "zh", false, w, h)
}
func CaptchaVerify(c *gin.Context, code string) bool {
	session := sessions.Default(c)
	if captchaId := session.Get("captcha"); captchaId != nil {
		session.Delete("captcha")
		_ = session.Save()
		if captcha.VerifyString(captchaId.(string), code) {
			return true
		} else {
			return false
		}
	} else {
		return false
	}
}
func Serve(w http.ResponseWriter, r *http.Request, id, ext, lang string, download bool, width, height int) error {
	w.Header().Set("Cache-Control", "no-cache, no-store, must-revalidate")
	w.Header().Set("Pragma", "no-cache")
	w.Header().Set("Expires", "0")

	var content bytes.Buffer
	switch ext {
	case ".png":
		w.Header().Set("Content-Type", "image/png")
		_ = captcha.WriteImage(&content, id, width, height)
	case ".wav":
		w.Header().Set("Content-Type", "audio/x-wav")
		_ = captcha.WriteAudio(&content, id, lang)
	default:
		return captcha.ErrNotFound
	}

	if download {
		w.Header().Set("Content-Type", "application/octet-stream")
	}
	http.ServeContent(w, r, id+ext, time.Time{}, bytes.NewReader(content.Bytes()))
	return nil
}

func main2() {
	router := gin.Default()
	router.LoadHTMLGlob("./*.html")
	router.Use(Session("topgoer"))
	router.GET("/captcha", func(c *gin.Context) {
		Captcha(c, 4)
	})
	router.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", nil)
	})
	router.GET("/captcha/verify/:value", func(c *gin.Context) {
		value := c.Param("value")
		if CaptchaVerify(c, value) {
			c.JSON(http.StatusOK, gin.H{"status": 0, "msg": "success"})
		} else {
			c.JSON(http.StatusOK, gin.H{"status": 1, "msg": "failed"})
		}
	})
	router.Run(":8080")
}

// token
//自定义一个字符串
var jwtkey = []byte("www.topgoer.com")
var str string

type Claims struct {
	UserId uint
	jwt.StandardClaims
}

func main() {
	r := gin.Default()
	r.GET("/set", setting)
	r.GET("/get", getting)
	//监听端口默认为8080
	r.Run(":8080")
}

//颁发token
func setting(ctx *gin.Context) {
	expireTime := time.Now().Add(7 * 24 * time.Hour)
	claims := &Claims{
		UserId: 2,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expireTime.Unix(), //过期时间
			IssuedAt:  time.Now().Unix(),
			Issuer:    "127.0.0.1",  // 签名颁发者
			Subject:   "user token", //签名主题
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	// fmt.Println(token)
	tokenString, err := token.SignedString(jwtkey)
	if err != nil {
		fmt.Println(err)
	}
	str = tokenString
	ctx.JSON(200, gin.H{"token": tokenString})
}

//解析token
func getting(ctx *gin.Context) {
	tokenString := ctx.GetHeader("Authorization")
	//vcalidate token formate
	if tokenString == "" {
		ctx.JSON(http.StatusUnauthorized, gin.H{"code": 401, "msg": "权限不足"})
		ctx.Abort()
		return
	}

	token, claims, err := ParseToken(tokenString)
	if err != nil || !token.Valid {
		ctx.JSON(http.StatusUnauthorized, gin.H{"code": 401, "msg": "权限不足"})
		ctx.Abort()
		return
	}
	fmt.Println(111)
	fmt.Println(claims.UserId)
}

func ParseToken(tokenString string) (*jwt.Token, *Claims, error) {
	Claims := &Claims{}
	token, err := jwt.ParseWithClaims(tokenString, Claims, func(token *jwt.Token) (i interface{}, err error) {
		return jwtkey, nil
	})
	return token, Claims, err
}
