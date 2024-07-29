package ctype

import "encoding/json"

// 角色类型，用于权限控制
type ArticleTypes int

const (
	PermissionLost       ArticleTypes = 1 //丢失的物品
	PermissionFound      ArticleTypes = 2 //寻找失主的物品
)

// 角色json序列化
func (r ArticleTypes) MarshalJSON() ([]byte, error) {
	return json.Marshal(r.String())
}

// 角色匹配
func (r ArticleTypes) String() string {
	switch r {
	case PermissionLost:
		return "寻找失物"
	case PermissionFound:
		return "寻找失主"
	default:
		return "其他"
	}
}
