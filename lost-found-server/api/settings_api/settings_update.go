package settings_api

import (
	"lost_found_server/config"
	"lost_found_server/core"
	"lost_found_server/global"
	"lost_found_server/models/res"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)
// SettingsInfoUpdateView 更新配置信息
func (SettingsApi) SettingsInfoUpdateView(c *gin.Context) {
	name := c.Param("name")
	
	switch name {
	case "qiniu":
		var si config.QiNiu
		//将请求体中json数据绑定到si结构体中
		err := c.ShouldBindJSON(&si)
		if err != nil {
			res.FailWithCode(res.ArgumentError, c)
			return
		}
		//配置信息修改
		global.Config.QiNiu = si
	
	case "upload":
		var si config.Upload
		//将请求体中json数据绑定到si结构体中
		err := c.ShouldBindJSON(&si)
		if err != nil {
			res.FailWithCode(res.ArgumentError, c)
			return
		}
		//配置信息修改
		global.Config.Upload = si
	default:
		res.FailWithMessage("没有对应的配置信息", c)
		return
	}

	//配置信息写入到yaml中(储存)
	err := core.SetYaml()
	if err != nil {
		//服务端查看
		logrus.Error(err)
		//返回客户端
		res.FailWithMessage(err.Error(), c)
		return
	}
	res.OkWithMessage("修改成功", c)
}
