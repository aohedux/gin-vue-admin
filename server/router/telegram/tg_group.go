package telegram

import (
	v1 "github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/gin-gonic/gin"
)

type TelegramGroupApi struct {
}

func (t *TelegramGroupApi) InitTelegramGroupRouter(Router *gin.RouterGroup) {

	TelegramGroupRouter := Router.Group("/api")
	api := v1.ApiGroupApp.TelegramApiGroup.TelegramGroupApi
	{
		TelegramGroupRouter.POST("/create_telegram_group", api.CreateTelegramGroupApi)
		TelegramGroupRouter.POST("/list_telegram_group", api.ListTelegramGroupApi)
		TelegramGroupRouter.POST("/update_telegram_group", api.UpdateTelegramGroupApi)
		TelegramGroupRouter.POST("/delete_telegram_group", api.DeleteTelegramGroupApi)
	}
}
