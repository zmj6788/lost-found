package models

import "lost_found_server/models/ctype"

// ArticlesModel 物品表
type ArticleModel struct {
	MODEL
	Name      string              `gorm:"size:36" json:"name"`                  //物品名称
	Place     string              `gorm:"size:36" json:"place"`                 //物品地点
	Time      string              `gorm:"size:36" json:"time"`                  //物品时间
	Picture   string              `gorm:"size:360" json:"picture"`               //物品图片
	Describe  string              `gorm:"size:360" json:"describe"`             //物品描述
	Types     ctype.ArticleTypes  `gorm:"type:tinyint" json:"types"`            //物品类型
	Status    ctype.ArticleStatus `gorm:"type:tinyint;default:1" json:"status"` //物品状态
	UserModel UserModel           `gorm:"foreignKey:UserID" json:"-"`           // 发布物品的作者
	UserID    uint                ` json:"user_id"`                             // 发布物品的用户id
}
