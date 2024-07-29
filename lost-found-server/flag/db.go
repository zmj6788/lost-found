package flag

import (
	"lost_found_server/global"
	"lost_found_server/models"

	"github.com/sirupsen/logrus"
)

func Makemigrations() {
	var err error
	//生成四张表的表结构
	err = global.DB.Set("gorm:table_options", "ENGINE=InnoDB").
		AutoMigrate(
			&models.BannerModel{},
			&models.UserModel{},
			&models.ArticleModel{},
		)
	if err != nil {
		logrus.Error("初始化数据库失败", err)
		return
	}
	logrus.Info("初始化数据库成功")
}
