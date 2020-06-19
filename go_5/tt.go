package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	// r.GET("/", func(c *gin.Context) {
	// 	c.String(http.StatusOK, "hello world")
	// })

	r.GET("/user/:name", func(c *gin.Context) {
		// 获取参数
		name := c.Param("name")
		// 返回浏览器
		c.String(http.StatusOK, name)
	})

	r.Run(":8888")
}
