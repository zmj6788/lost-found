package users_api

import (
	"lost_found_server/global"
	"lost_found_server/models"
	"lost_found_server/models/res"

	"github.com/gin-gonic/gin"
)

type RegisterRequest struct {
	UserName string `json:"user_name"` // 文件名
	Password string `json:"password"`  // 是否上传成功
}
//简易版注册
func (UsersApi) RegisterView(c *gin.Context) {
	var newUser RegisterRequest
	err := c.ShouldBindJSON(&newUser)
	if err != nil {
		res.FailWithCode(res.ArgumentError, c)
		return
	}
	if newUser.UserName == "" || newUser.Password == "" {
		res.FailWithMessage("用户名或密码不能为空", c)
		return
	} 
	//获取参数成功，验证用户名是否已存在
	if global.DB.Where("user_name=?", newUser.UserName).Find(&models.UserModel{}).RowsAffected > 0 {
		res.FailWithMessage("用户名已存在", c)
		return
	}
	//验证通过，创建用户
	var user models.UserModel
	user.UserName = newUser.UserName
	user.Password = newUser.Password
	user.Avatar = "https://q.qlogo.cn/g?b=qq&nk=2124427385&s=100"
	global.DB.Create(&user)
	res.OkWithMessage("注册成功", c)
}
