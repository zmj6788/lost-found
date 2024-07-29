package images_api

import (
	"fmt"
	"lost_found_server/global"
	"lost_found_server/models"
	"lost_found_server/models/res"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

type ImagesRemoveResponse struct {
	ID        string `json:"id"`         // 文件名
	IsSuccess bool   `json:"is_success"` // 是否上传成功
	Msg       string `json:"msg"`        // 提示信息

}

// 图片删除接口，接受一个删除请求，请求参数为idlist的json数据
func (ImagesApi) ImageRemoveView(c *gin.Context) {

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

	//获得删除请求参数信息
	var rq models.RemoveRequest
	err = c.ShouldBindJSON(&rq)
	if err != nil {
		res.FailWithCode(res.ArgumentError, c)
		logrus.Info(err)
		return
	}
	var resList []ImagesRemoveResponse
	var a int64
	//rq.IDList是我们传入要删除的图片id，在删除前判断是否有权限删除
	// 遍历物品ID数组
	for _, itemId := range rq.IDList {
		var item models.BannerModel
		var response ImagesRemoveResponse
		// 查询物品以获取其用户ID
		if err := global.DB.Where("id = ?", itemId).First(&item).Error; err != nil {
			response.ID = strconv.Itoa(int(itemId))
			response.IsSuccess = false
			response.Msg = "文件不存在"
			// 如果查询失败，可能是物品不存在
			resList = append(resList, response)
			continue
		}
		// 检查物品的用户ID是否与输入的用户ID匹配
		if item.UserID != uintID {
			// 不匹配，提前返回false
			response.ID = strconv.Itoa(int(itemId))
			response.IsSuccess = false
			response.Msg = "你没有权限删除此图片"
			// 如果查询失败，可能是物品不存在
			resList = append(resList, response)
			continue
		}
		//这是要删除的id
		var imagelist []models.BannerModel
		//获取要删除的数据到imagelist中
		global.DB.Find(&imagelist, itemId)
		response.ID = strconv.Itoa(int(itemId))
			response.IsSuccess = true
			response.Msg = "成功删除"
			// 如果查询失败，可能是物品不存在
			resList = append(resList, response)
		a ++
		//imagelist中既是需要删除的图片数据
		global.DB.Delete(&imagelist)
		logrus.Info(imagelist)
	}
	res.Ok(resList, fmt.Sprintf("共删除%d张图片", a), c)
}
