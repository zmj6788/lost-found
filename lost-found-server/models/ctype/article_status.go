package ctype

import "encoding/json"

// 角色类型，用于权限控制
type ArticleStatus int

const (
	PermissionOk       ArticleStatus = 1 //未找回
	PermissionNo      ArticleStatus = 2 //已找回
)

// 角色json序列化
func (r ArticleStatus) MarshalJSON() ([]byte, error) {
	return json.Marshal(r.String())
}

// 角色匹配
func (r ArticleStatus) String() string {
	switch r {
	case PermissionOk:
		return "未找到"
	case PermissionNo:
		return "已找到"
	default:
		return "其他"
	}
}
