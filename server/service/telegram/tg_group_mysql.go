package telegram

import (
	"fmt"

	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/telegram"
	"github.com/flipped-aurora/gin-vue-admin/server/model/telegram/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/telegram/response"
)

type TelegramGroupApi struct {
}

// 操作前判断是否重复
func (t *TelegramGroupApi) CheckRepeat(groupName string, userName string) (err error) {

	var count int64
	err = global.GVA_DB.Model(&telegram.TelegramGroup{}).Where("group_name = ? and user_name = ?", groupName, userName).Count(&count).Error
	if count > 0 {
		return fmt.Errorf("该群组已存在")
	}
	return nil
}

func (t *TelegramGroupApi) AddTelegramGroup(groupName string, userName string) (err error) {

	err = t.CheckRepeat(groupName, userName)
	if err != nil {
		return err
	}
	err = global.GVA_DB.Create(&telegram.TelegramGroup{GroupName: groupName, UserName: userName}).Error

	if err != nil {
		return err
	}
	return nil
}
func (t *TelegramGroupApi) ListTelegramGroup(page int, pagenum int, username string, groupname string) (response.ListTelegramGroupResponse, error) {
	var count int64
	var list response.ListTelegramGroupResponse
	global.GVA_DB.Model(&telegram.TelegramGroup{}).Where("user_name = ? AND group_name LIKE ? ", username, "%"+groupname+"%").Offset((page - 1) * pagenum).Limit(pagenum).Find(&list.List)
	global.GVA_DB.Model(&telegram.TelegramGroup{}).Where("user_name = ?", username).Count(&count)
	fmt.Println("总条数：", count)

	list.Total = count

	return list, nil
}

func (t *TelegramGroupApi) UpdateTelegramGroup(req request.UpdateTelegramGroupRequest) error {
	errCheck := t.CheckRepeat(req.GroupName, req.UserName)
	fmt.Println(req.GroupName, req.UserName)
	if errCheck != nil {
		return errCheck
	}
	err := global.GVA_DB.Model(&telegram.TelegramGroup{}).Where("id = ?", req.Id).Update("group_name", req.GroupName).Error
	if err != nil {
		return err
	}
	return nil
}
func (t *TelegramGroupApi) DeleteTelegramGroup(req request.DeleteTelegramGroupRequest) error {

	err := global.GVA_DB.Model(&telegram.TelegramGroup{}).Where("id in ?", req.Id).Delete(&telegram.TelegramGroup{}).Error
	if err != nil {
		return err
	}
	return nil
}
