package common

import (
	"lost_found_server/global"
	"lost_found_server/models"

	"gorm.io/gorm"
)

type Option struct {
	models.PageInfo
	Debug bool //是否开启debug模式,是否查看日志
}

// ComList 通用列表分页查询服务
func ComList[T any](userID uint, model T, option Option) (list []T, count int64, err error) {

	DB := global.DB
	if option.Debug {
		DB = global.DB.Session(&gorm.Session{Logger: global.MysqlLog})
	}

	//排序
	if option.Sort == "" {
		option.Sort = "created_at desc" //默认排序创建顺序desc从晚到早.asc从早到晚
	}

	count = DB.Where("user_id = ?", userID).Select("id").Find(&list).RowsAffected
	
	//偏移量
	offset := (option.Page - 1) * option.Limit
	//如果偏移量小于0，则从0开始
	if offset < 0 {
		offset = 0
	}
	//如果limit为0，则查询所有
	if option.Limit == 0 {
		option.Limit = -1
	}
	err = DB.Where("user_id = ? and name LIKE ?", userID,"%"+option.Key+"%").Limit(option.Limit).Offset(offset).Order(option.Sort).Find(&list).Error
	
	if option.ID != 0 {
		err = DB.Where("id = ?", option.ID).First(&list).Error
	}
	
	return list, count, err
}
