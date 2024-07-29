package routers

import (
	"lost_found_server/api"

	"github.com/gin-gonic/gin"
)

func ImagesRouter(router *gin.RouterGroup) {
	
	ImagesApi := api.ApiGroupApp.ImagesApi
	router.GET("/images/:userid", ImagesApi.ImageListView)
	router.POST("/images/:userid", ImagesApi.ImageUploadView)
	router.DELETE("/images/:userid", ImagesApi.ImageRemoveView)
	router.PUT("/images/:userid", ImagesApi.ImageUpdateView)
	//静态文件挂载
	//访问示例
	//http://localhost:8080/api/uploads/file/1.jpg
	router.Static("/uploads/file", "./uploads/file")
}
