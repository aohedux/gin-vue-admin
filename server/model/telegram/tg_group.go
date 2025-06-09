package telegram

import "github.com/flipped-aurora/gin-vue-admin/server/global"

type TelegramGroup struct {
	global.GVA_MODEL
	UserName  string `json:"userName" gorm:"comment:用户登录名"`
	GroupName string `json:"groupName" gorm:"comment:"`
}

func (TelegramGroup) TableName() string {
	return "tg_group"
}
