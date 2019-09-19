package main

import (
	"crypto/sha1"
	"fmt"
	"github.com/gin-gonic/gin"
	"io"
	"log"
	"net/http"
	"sort"
	"strings"
)

const (
	token = "hlc944792" //设置token
)

func main() {
	// 初始化引擎
	engine := gin.Default()
	// 注册一个路由和处理函数
	engine.Any("/", WebRoot)

	// 注册一个动态路由
	// 可以匹配 /user/joy
	// 不能匹配 /user 和 /user/
	engine.GET("/user/:name", func(c *gin.Context) {
		// 使用 c.Param(key) 获取 url 参数
		name := c.Param("name")
		c.String(http.StatusOK, "Hello %s", name)
	})

	// 注册一个高级的动态路由
	// 该路由会匹配 /user/john/ 和 /user/john/send
	// 如果没有任何路由匹配到 /user/john, 那么他就会重定向到 /user/john/，从而被该方法匹配到
	engine.GET("/user/:name/*action", func(c *gin.Context) {
		name := c.Param("name")
		action := c.Param("action")
		message := name + " is " + action
		c.String(http.StatusOK, message)
	})
	// 绑定端口，然后启动应用
	engine.Run(":80")
}

/**
* 根请求处理函数
* 所有本次请求相关的方法都在 context 中，完美
* 输出响应 hello, world
 */
func WebRoot(context *gin.Context) {

	timestamp := context.Param("timestamp")
	nonce := context.Param("nonce")
	signature := context.Param("signature")
	echostr := context.Param("echostr")

	log.Println("timestamp "+timestamp)
	log.Println("nonce "+nonce)
	log.Println("signature "+signature)
	log.Println("echostr "+echostr)
	log.Println("echostr "+echostr)

	si := []string{token, timestamp, nonce}
	sort.Strings(si)            //字典序排序
	str := strings.Join(si, "") //组合字符串
	s := sha1.New()             //返回一个新的使用SHA1校验的hash.Hash接口
	io.WriteString(s, str)      //WriteString函数将字符串数组str中的内容写入到s中
	hashcode := fmt.Sprintf("%x", s.Sum(nil))

	log.Println("hashcode "+hashcode)

	if hashcode != signature {
		log.Println("Wechat Service: This http request is not from wechat platform")
		return
	}
	log.Println("validateUrl Ok")
	log.Println("validateUrl Ok")
	context.String(http.StatusOK,echostr)
}
