package articles_api

import (
	"lost_found_server/global"
	"lost_found_server/models"
	"lost_found_server/models/res"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

func (ArticleApi) ArticleCreateView(c *gin.Context) {

	//获取一个用户id，用于将物品关联到用户
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
		res.FailWithMessage("用户ID不存在", c)
		return
	}
	//uintID发布物品的用户的id
	var si models.ArticleModel
	err = c.ShouldBindJSON(&si)
	si.UserID = uintID
	if err != nil {
		res.FailWithCode(res.ArgumentError, c)
		return
	}
	//参数匹配成功
	if err := global.DB.Create(&si).Error; err != nil {
		logrus.Error(err)
		res.FailWithMessage("添加物品失败", c)
		return
	}
	
	res.OkWithMessage("添加物品成功", c)
}
