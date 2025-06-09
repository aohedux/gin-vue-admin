package telegram

import "github.com/flipped-aurora/gin-vue-admin/server/service"

type ApiGroup struct {
	TelegramMyApi
	TelegramGroupApi
	TelegramContactListApi
}

var (
	telegramService            = service.ServiceGroupApp.TelegramServiceGroup.TelegramGroupApi
	telegramMyService          = service.ServiceGroupApp.TelegramServiceGroup.TelegramMyApi
	telegramContactListService = service.ServiceGroupApp.TelegramServiceGroup.TelegramContactListApi
)
