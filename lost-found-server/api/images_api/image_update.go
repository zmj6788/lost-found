package images_api

import (
	"lost_found_server/global"
	"lost_found_server/models"
	"lost_found_server/models/res"
	"strconv"

	"github.com/gin-gonic/gin"
)

// 图片更新请求
type ImageUpdateRequest struct {
	ID   uint   `json:"id" binding:"required" msg:"请选择有效文件id"`
	Name string `json:"name" binding:"required" msg:"请输入修改后的文件名称"`
}

// 图片更新接口函数
func (ImagesApi) ImageUpdateView(c *gin.Context) {

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

	var req ImageUpdateRequest
	err = c.ShouldBindJSON(&req)
	if err != nil {
		//响应提示信息
		res.FailWithError(err, &req, c)
		return
	}
	var imageModel models.BannerModel
	err = global.DB.Take(&imageModel, req.ID).Error
	if err != nil {
		res.FailWithMessage("图片不存在", c)
		return
	}
	if imageModel.UserID != uintID {
		res.FailWithMessage("图片不属于该用户", c)
		return
	}
	// global.DB.Model(&imageModel).Updates(models.BannerModel{
	// 	Name: req.Name,
	// })
	//修改表数据的两种方法
	err = global.DB.Model(&imageModel).Update("name", req.Name).Error
	if err != nil {
		res.FailWithMessage(err.Error(), c)
		return
	}
	res.OkWithMessage("图片名称修改成功", c)
}
