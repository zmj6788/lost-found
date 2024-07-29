package users_api

import (
	"lost_found_server/global"
	"lost_found_server/models"
	"lost_found_server/models/res"

	"github.com/gin-gonic/gin"
)

type LoginRequest struct {
	UserName string `json:"user_name"` // 文件名
	Password string `json:"password"`  // 是否上传成功
}

func (UsersApi) LoginView(c *gin.Context) {
	var newUser LoginRequest
	err := c.ShouldBindJSON(&newUser)
	if err != nil {
		res.FailWithCode(res.ArgumentError, c)
		return
	}
	if newUser.UserName == "" || newUser.Password == "" {
		res.FailWithMessage("用户名或密码不能为空", c)
		return
	}
	//获取参数成功，验证用户名与密码是否正确
	if global.DB.Where("user_name=? and password=?", newUser.UserName, newUser.Password).Find(&models.UserModel{}).RowsAffected < 1 {
		res.FailWithMessage("用户名或密码错误", c)
		return
	}
	var user models.UserModel
	global.DB.Where("user_name = ?", newUser.UserName).Find(&user)
	if user.Role == 1 {
		res.Ok(gin.H{
		"userid": user.ID,
	}, "管理员登录", c)
		return
	}
	res.Ok(gin.H{
		"userid": user.ID,
	}, "登录成功", c)
}
