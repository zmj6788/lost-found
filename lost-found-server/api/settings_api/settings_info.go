package settings_api

import (
	"lost_found_server/global"
	"lost_found_server/models/res"

	"github.com/gin-gonic/gin"
)

type SettingsUri struct {
	Name string `uri:"name"`
}
// SettingsInfoView 获取配置信息
func (SettingsApi) SettingsInfoView(c *gin.Context) {
	var uri SettingsUri
	// 绑定uri参数
	err := c.ShouldBindUri(&uri)
	if err != nil {
		res.FailWithMessage(err.Error(), c)
		return
	}
	name := uri.Name
	//name := c.Param("name")
	switch name {
	case "qiniu":
		res.OkWithData(global.Config.QiNiu, c)
	
	case "upload":
		res.OkWithData(global.Config.Upload, c)
	default:
		res.FailWithMessage("没有对应的配置信息", c)
	}
}
