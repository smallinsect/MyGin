package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main1() {
	r := gin.Default()

	r.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "hello world")
	})

	r.GET("/user/:name", func(c *gin.Context) {
		// 获取参数
		name := c.Param("name")
		// 返回浏览器
		c.String(http.StatusOK, name)
	})

	r.Run(":8888")
}

func main() {
	r := gin.Default()
	// 多种相应方式
	// 1.json
	r.GET("/someJSON", func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "someJSON", "status": 200})
	})

	// 2.结构体
	r.GET("/someStruct", func(c *gin.Context) {
		var msg struct {
			Name    string
			Message string
			Number  int64
		}

		msg.Name = "小昆虫"
		msg.Message = "dafdsaf"
		msg.Number = 1111
		c.JSON(200, msg)
	})

	// 3.XML
	r.GET("/someXML", func(c *gin.Context) {
		c.XML(200, gin.H{"message": "abc"})
	})

	// 4.YAML
	r.GET("/someYAML", func(c *gin.Context) {
		c.YAML(200, gin.H{"message": "CCCCCC"})
	})

	// 5.protoBuf格式，谷歌开发的高效存储读取的工具
	// r.GET("/someProtoBuf", func(c *gin.Context) {
	// 	reps := []int64{int64(1), int64(2)}
	// 	label := "label"

	// 	data := &protoexample.Test{
	// 		Label: &label,
	// 		Reps:  reps,
	// 	}

	// 	c.ProtoBuf(200, data)
	// })

	r.Run(":8888")
}
