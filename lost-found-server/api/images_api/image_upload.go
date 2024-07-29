package images_api

import (
	"io/fs"
	"lost_found_server/global"
	"lost_found_server/models"
	"lost_found_server/models/res"
	"lost_found_server/service"
	"lost_found_server/service/image_service"
	"os"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

var (
	// 规定图片上传白名单
	WhiteImageList = []string{
		"pjp",
		"svgz",
		"jpg",
		"jpeg",
		"ico",
		"tiff",
		"gif",
		"svg",
		"jfif",
		"webp",
		"png",
		"bmp",
		"pjpeg",
		"avif",
	}
)

// 规定文件上传的响应格式，便于客户端解析上传结果信息
type FileUploadResponse struct {
	FileName  string `json:"file_name"`  // 文件名
	IsSuccess bool   `json:"is_success"` // 是否上传成功
	Msg       string `json:"msg"`        // 提示信息
}

// ImageUploodView上传到个图片，返回图片的url
func (ImagesApi) ImageUploadView(c *gin.Context) {
	//再次出获取一个用户id，用于将图片关联到用户
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
	
//判断用户id是否存在	
	var count int64
	err = global.DB.Model(&models.UserModel{}).Where("id = ?", uintID).Count(&count).Error
	if err != nil {
		res.FailWithError(err, "系统错误" , c)
		return
	}
	if count == 0 {
		res.FailWithMessage("用户ID不存在", c)
		return
	}

	//获取单个上传的文件
	//FileHeader, err := c.FormFile("image")

	//获取多个上传的文件
	form, err := c.MultipartForm()
	if err != nil {
		res.FailWithMessage(err.Error(), c)
	}
	fileList, ok := form.File["images"]
	if !ok {
		res.FailWithMessage("文件不存在", c)
		return
	}
	//判断图片文件存储路径是否存在
	//不存在创建
	basepath := global.Config.Upload.Path
	_, err = os.ReadDir(basepath)
	if err != nil {
		err := os.MkdirAll(basepath, fs.ModePerm)
		if err != nil {
			res.FailWithMessage(err.Error(), c)
			logrus.Error(err)
			return
		}
	}
	//遍历文件，做上传文件操作
	var resList []image_service.FileUploadResponse
	for _, file := range fileList {

		// 上传文件
		serviceRes := service.Services.ImageService.ImageUploadService(file, uintID)
		if !serviceRes.IsSuccess {
			resList = append(resList, serviceRes)
			continue
		}
		// 成功的
		if !global.Config.QiNiu.Enable {
			// 本地还得保存一下
			//上传成功但是七牛云存储没有开启，说明是本地存储成功，保存一下
			err = c.SaveUploadedFile(file, serviceRes.FileName)
			if err != nil {
				global.Log.Error(err)
				serviceRes.Msg = err.Error()
				serviceRes.IsSuccess = false
				resList = append(resList, serviceRes)
				continue
			}
		}
		resList = append(resList, serviceRes)
	}

	res.Ok(resList, "响应成功", c)

}
