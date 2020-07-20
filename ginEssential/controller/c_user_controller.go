package controller

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"xuqiulin.com/mygin/ginEssential/common"
	"xuqiulin.com/mygin/ginEssential/model"
)

func Register(ctx *gin.Context) {

	DB := common.GetDB()

	// 获取参数
	name := ctx.PostForm("name")
	telephone := ctx.PostForm("telephone")
	password := ctx.PostForm("password")

	// 数据验证
	if len(telephone) != 11 {
		ctx.JSON(http.StatusUnprocessableEntity, gin.H{
			"code": 422, "msg": "手机号没有11位",
		})
	}
	if len(password) < 6 {
		ctx.JSON(http.StatusUnprocessableEntity, gin.H{
			"code": 422, "msg": "密码少于6位",
		})
	}
	if len(name) == 0 {
		ctx.JSON(http.StatusUnprocessableEntity, gin.H{
			"code": 422, "msg": "名字长度不能为0",
		})
	}

	log.Println(name, telephone, password)

	newUser := model.User{
		Name:      name,
		Telephone: telephone,
		Password:  password,
	}
	DB.Create(&newUser)

	ctx.JSON(200, gin.H{
		"msg": "pong",
	})
}

func Login(ctx *gin.Context) {
	DB := common.GetDB()

	name := ctx.PostForm("name")
	password := ctx.PostForm("password")

	user := model.User{}
	DB.Where("name = ? and password = ?", name, password).First(&user)
	if user.ID == 0 {
		ctx.JSON(http.StatusUnprocessableEntity,
			gin.H{
				"code": 422,
				"msg":  "用户不存在",
			},
		)
		return
	}
	// bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)

	token, err := common.ReleaseToken(&user)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError,
			gin.H{
				"code": 500,
				"msg":  "系统错误",
			},
		)
		log.Println("token generate error", err)
		return
	}
	// 用户存在 发放token
	ctx.JSON(200, gin.H{
		"code": 200,
		"data": gin.H{"token": token},
		"msg":  "登录成功",
	})
}

func Info(ctx *gin.Context) {
	// 直接从上下文获取信息
	user, _ := ctx.Get("user")
	ctx.JSON(http.StatusOK, gin.H{"code": 200, "data": gin.H{"user": user}})
}
