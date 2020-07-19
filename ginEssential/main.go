package main

import (
	"xuqiulin.com/mygin/ginEssential/common"

	"github.com/gin-gonic/gin"
)

func main() {

	//初始化数据库
	db := common.InitDB()
	defer db.Close()

	r := gin.Default()

	r = CollectRouter(r)

	r.Run() // listen and serve on 0.0.0.0:8080
}
