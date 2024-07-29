package users_api

import (
	"lost_found_server/global"
	"lost_found_server/models"
	"lost_found_server/models/res"
	"strconv"

	"github.com/gin-gonic/gin"
)

// 用户查询个人信息
func (UsersApi) UserReadView(c *gin.Context) {

	//获取一个用户id，用于匹配对应的用户的图片列表
	intID, err := strconv.Atoi(c.Param("userid"))
	if err != nil {
		// 处理错误，例如用户ID不是一个有效的数字
		res.FailWithMessage("用户ID格式不正确", c)
		return
	}
	var uintID uint
	if intID < 0 {
		// 如果可能为负，需要处理这种情况，因为uint不能表示负数
		res.FailWithMessage("用户ID不能为负数", c)
		return
	}
	//转换id格式合法
	uintID = uint(intID)
	//判断用户id是否存在
	var count int64
	err = global.DB.Model(&models.UserModel{}).Where("id = ?", uintID).Count(&count).Error
	if err != nil {
		res.FailWithError(err, "系统错误", c)
		return
	}
	if count == 0 {
		res.FailWithMessage("用户不存在", c)
		return
	}
	
	//根据用户id查询用户信息
	var UserModel models.UserModel

	err = global.DB.Take(&UserModel, uintID).Error
	if err != nil {
		res.FailWithError(err, "系统错误", c)
		return
	}
	
	
	res.Ok(UserModel, "查询成功", c)

}
