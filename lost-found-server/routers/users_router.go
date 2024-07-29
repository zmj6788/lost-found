package routers

import (
	"lost_found_server/api"

	"github.com/gin-gonic/gin"
)

func UsersRouter(router *gin.RouterGroup) {
	UsersApi := api.ApiGroupApp.UsersApi
	router.POST("/register", UsersApi.RegisterView)
	router.POST("/login", UsersApi.LoginView)
	//用户信息更新
	router.PUT("/user/:userid", UsersApi.UserUpdateView)
	//注销账号
	router.DELETE("/user/:userid", UsersApi.UserDeleteView)
	//获取用户信息
	router.GET("/user/:userid", UsersApi.UserReadView)
}
