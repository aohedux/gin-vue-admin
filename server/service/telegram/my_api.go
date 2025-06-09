package telegram

import (
	"fmt"

	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
	"github.com/flipped-aurora/gin-vue-admin/server/model/telegram"
)

type TelegramMyApi struct {
}

func (t *TelegramMyApi) ListTelegramUser(req response.GetSessionResponse) ([]response.Session2, error) {

	result := make([]response.Session2, 0, len(req.Data.Session))

	for _, v := range req.Data.Session {
		var gour_name telegram.TelegramGroup
		err := global.GVA_DB.Model(&telegram.TelegramGroup{}).Where("id = ?", v.Tag).Find(&gour_name).Error
		if err != nil {
			return result, err

		}
		fmt.Println("分组名:", gour_name.GroupName)
		dto := response.Session2{
			Phone:     v.Phone,
			Username:  v.Username,
			Name:      v.Name,
			IsPremium: v.IsPremium,
			IsBlocked: v.IsBlocked,
			Online:    v.Online,
			Server:    v.Server,
			Proxy:     v.Proxy,
			CreatedAt: v.CreatedAt,
			Tag:       gour_name.GroupName,
		}
		result = append(result, dto)
	}
	return result, nil
}
