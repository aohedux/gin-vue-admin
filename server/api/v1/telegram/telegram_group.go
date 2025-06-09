package telegram

import (
	"fmt"

	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
	"github.com/flipped-aurora/gin-vue-admin/server/model/telegram/request"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type TelegramGroupApi struct {
}

func (t *TelegramGroupApi) CreateTelegramGroupApi(c *gin.Context) {
	var telegramGroupRequest request.TelegramGroupRequest
	err := c.ShouldBindJSON(&telegramGroupRequest)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	fmt.Println("创建群组:", telegramGroupRequest)
	errMysql := telegramService.AddTelegramGroup(telegramGroupRequest.GroupName, telegramGroupRequest.UserName)
	if errMysql != nil {

		response.Result(404, nil, errMysql.Error(), c)
		return
	}
	response.Ok(c)

}
func (t *TelegramGroupApi) ListTelegramGroupApi(c *gin.Context) {
	var listRequest request.ListTelegramGroupRequest
	err := c.ShouldBindJSON(&listRequest)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	list, errlist := telegramService.ListTelegramGroup(listRequest.Page, listRequest.PageSize, listRequest.UserName, listRequest.GroupName)
	if errlist != nil {
		return
	}

	response.Result(0, list, "加载成功", c)
}
func (t *TelegramGroupApi) UpdateTelegramGroupApi(c *gin.Context) {
	var updateRequest request.UpdateTelegramGroupRequest

	err := c.ShouldBindJSON(&updateRequest)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = telegramService.UpdateTelegramGroup(updateRequest)
	if err != nil {
		response.FailWithMessage("更新失败:"+err.Error(), c)
		return
	}
	response.OkWithMessage("更新成功", c)

}

func (t *TelegramGroupApi) DeleteTelegramGroupApi(c *gin.Context) {

	var deleteRequest request.DeleteTelegramGroupRequest
	err := c.ShouldBindJSON(&deleteRequest)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = telegramService.DeleteTelegramGroup(deleteRequest)
	if err != nil {
		global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败:"+err.Error(), c)
		return
	}
	response.OkWithMessage("删除成功", c)
}
