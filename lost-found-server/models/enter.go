package models

import "time"

// 自定义MODEL，没有用gorm的MODEL,因为我们不需要逻辑删除
type MODEL struct {
	ID        uint      `gorm:"primarykey" json:"id"` // 主键ID
	CreatedAt time.Time `json:"created_at"`           // 创建时间
	UpdatedAt time.Time `json:"-"`                    // 更新时间
}

// 用于分页查询
type PageInfo struct {
	Page  int    `form:"page"`  //页码
	Key   string `form:"key"`   //搜索关键字
	Limit int    `form:"limit"` //每页显示多少条
	Sort  string `form:"sort"`  //排序
	ID    uint   `form:"id"`    //id查找
}

//用于接收删除请求
type RemoveRequest struct {
	IDList []uint `json:"id_list"`
}