package telegram

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	resp1 "github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
	"github.com/flipped-aurora/gin-vue-admin/server/model/telegram/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/telegram/response"

	"github.com/gin-gonic/gin"
)

type TelegramContactListApi struct {
}

func (t *TelegramContactListApi) GetTelegramContactList(c *gin.Context) {
	var req request.ListTelegramContactlistRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	reqlist, err := json.Marshal(&req)
	if err != nil {
		panic(err)
	}
	fmt.Println(bytes.NewBuffer(reqlist))
	resp, err := http.Post("http://localhost:9090/api/get_contacts", "application/json", bytes.NewBuffer(reqlist))
	if err != nil {
		panic(err)
	}

	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	var result response.ListTelegramContactListResponse

	if err := json.Unmarshal(body, &result); err != nil {
		panic(err)
	}
	fmt.Println(string(body))
	fmt.Println(result)
	resp1.Result(0, result, "加载成功", c)

}
