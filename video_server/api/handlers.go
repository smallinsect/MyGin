package main

import (
	"io"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

// CreateUser 创建用户
func CreateUser(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	io.WriteString(w, "Create User Handler")
}

// Login 用户登录
func Login(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	uname := ps.ByName("user_name")
	io.WriteString(w, uname)
}
