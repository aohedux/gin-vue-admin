package telegram

import (
	v1 "github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/gin-gonic/gin"
)

type TelegramMyApi struct {
}

func (m *TelegramMyApi) InitTelegramRouter(Router *gin.RouterGroup) {

	TelegramGroup := Router.Group("/api")
	api := v1.ApiGroupApp.TelegramApiGroup.TelegramMyApi
	{
		TelegramGroup.POST("/get_session", api.TelegramApi)
		TelegramGroup.POST("/import_session", api.AddTelegramApi)
		TelegramGroup.POST("/delete_session", api.DeleteTelegramApi)
		TelegramGroup.POST("/sign_in", api.LoginTelegramApi)
		TelegramGroup.POST("/sign_out", api.OutLoginTelegramApi)
		TelegramGroup.POST("/change_tag", api.ChangeTagTelegramApi)
	}
}

// func InitRouter(Router *gin.RouterGroup) {

// 	api := Router.Group("/api")
// 	{
// api.POST("/get_session", system.HandlerGetSession)
// 		// api.POST("/import_session", client.ImportSession)
// 		// api.POST("/delete_session", client.DeleteSession)
// 		// api.POST("/sign_in", client.SignIn)
// 		// api.POST("/sign_out", client.SignOut)
// 	}

// }
