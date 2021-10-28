package main

import "github.com/gin-gonic/gin"

func main() {
	r := gin.Default() // 会初始化时导入logger 和 recovery 中间件： 输出日志和异常捕获
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "pong"})
	})
	r.Run()
}
