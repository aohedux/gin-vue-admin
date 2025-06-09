package telegram

import (
	v1 "github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/gin-gonic/gin"
)

type TelegramContactListApi struct {
}

func (t *TelegramContactListApi) InitTelegramContactListRouter(Router *gin.RouterGroup) {
	TelegramContactListRouter := Router.Group("/api")
	api := v1.ApiGroupApp.TelegramApiGroup.TelegramContactListApi
	{
		TelegramContactListRouter.POST("/get_contacts", api.GetTelegramContactList)

	}
}
