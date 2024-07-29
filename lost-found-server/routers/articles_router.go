package routers

import (
	"lost_found_server/api"
	"github.com/gin-gonic/gin"
)

func ArticlesRouter(router *gin.RouterGroup) {
	ArticleApi := api.ApiGroupApp.ArticleApi
	router.POST("/articles/:userid", ArticleApi.ArticleCreateView)
	router.GET("/articles/:userid", ArticleApi.ArticleReadView)
	router.PUT("/articles/:userid", ArticleApi.ArticleUpdateView)
	router.DELETE("/articles/:userid", ArticleApi.ArticleDeleteView)
}
