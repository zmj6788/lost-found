package articles_api

import (
	"lost_found_server/global"
	"lost_found_server/models"
	"lost_found_server/models/res"
	"strconv"

	"github.com/gin-gonic/gin"
)

func (ArticleApi) ArticleUpdateView(c *gin.Context) {

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
		res.FailWithMessage("用户ID不存在", c)
		return
	}
	var ArticleModel models.ArticleModel
	err = c.ShouldBindJSON(&ArticleModel)
	if err != nil {
		//响应提示信息
		res.FailWithError(err, &ArticleModel, c)
		return
	}
	var before models.ArticleModel
	ArticleModel.UserID = uintID

	err = global.DB.Take(&before, ArticleModel.ID).Error
	if err != nil {
		res.FailWithMessage("物品不存在", c)
		return
	}
	if before.UserID != uintID {
		res.FailWithMessage("物品不属于该用户", c)
		return
	}
	// global.DB.Model(&imageModel).Updates(models.BannerModel{
	// 	Name: req.Name,
	// })
	//修改表数据的两种方法
	
	err = global.DB.Model(&before).Updates(&ArticleModel).Error
	if err != nil {
		res.FailWithMessage(err.Error(), c)
		return
	}
	res.OkWithMessage("物品修改成功", c)
}
