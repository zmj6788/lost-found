package api

import (
	images_api "lost_found_server/api/images_api"
	settings_api "lost_found_server/api/settings_api"
	users_api "lost_found_server/api/users_api"
	articles_api "lost_found_server/api/articles_api"
)

type ApiGroup struct {
	SettingsApi settings_api.SettingsApi
	ImagesApi   images_api.ImagesApi
	UsersApi    users_api.UsersApi
	ArticleApi  articles_api.ArticleApi
}

var ApiGroupApp = new(ApiGroup)
