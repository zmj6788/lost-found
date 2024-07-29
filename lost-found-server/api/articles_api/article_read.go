package articles_api

import (
	"lost_found_server/global"
	"lost_found_server/models"
	"lost_found_server/models/res"
	"lost_found_server/service/common"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

func (ArticleApi) ArticleReadView(c *gin.Context) {

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
	var page models.PageInfo
	err = c.ShouldBindQuery(&page)
	logrus.Info(page)
	if err != nil {
		res.FailWithCode(res.ArgumentError, c)
		return
	}
	//使用封装的列表分页查询服务
	list, count, err := common.ComList(
		uintID,
		models.ArticleModel{},
		common.Option{
			PageInfo: page,
			Debug:    true,
		})
	if err != nil {
		res.FailWithMessage("物品列表获取失败", c)
		return
	}
	if len(list) == 0 { // 更改为检查list的长度，因为这才是最终查询的结果
		res.OkWithMessage("满足查询条件的物品列表为空", c)
		return
	}
	res.OkWithList(list, count, c)
}
