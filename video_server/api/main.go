package main

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

// RegisterHandlers 注册
func RegisterHandlers() *httprouter.Router {
	router := httprouter.New()

	router.POST("/user", CreateUser)
	router.POST("/user/:user_name", Login) // post url:localhost:8000/user/1

	return router
}

func main() {
	r := RegisterHandlers()
	http.ListenAndServe(":8000", r)
}
