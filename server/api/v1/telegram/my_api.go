package telegram

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"

	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
	"github.com/gin-gonic/gin"
)

type TelegramMyApi struct {
}

// 响应结构体定义

func (m *TelegramMyApi) TelegramApi(c *gin.Context) {
	var getSessionRequest request.GetSessionRequest
	var listResult response.GetSessionResponseTwo
	if err := c.ShouldBindJSON(&getSessionRequest); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	req, err := json.Marshal(&getSessionRequest)
	if err != nil {
		panic(err)
	}
	fmt.Println(bytes.NewBuffer(req))
	resp, err := http.Post("http://localhost:9090/api/get_session", "application/json", bytes.NewBuffer(req))
	if err != nil {
		panic(err)
	}

	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)

	var result response.GetSessionResponse
	if err := json.Unmarshal(body, &result); err != nil {
		panic(err)
	}

	for k, v := range result.Data.Session {
		fmt.Println(k, v)
	}
	listResp, err := telegramMyService.ListTelegramUser(result)
	if err != nil {
		panic(err)
	}

	listResult.Session2 = listResp
	listResult.Count = result.Data.Count
	response.Result(0, listResult, "加载成功", c)
}

func (m *TelegramMyApi) AddTelegramApi(c *gin.Context) {

	var req request.ImportSessionRequest

	// 解析表单
	if err := c.ShouldBind(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	// 创建管道和multipart写入器
	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)
	// 添加用户名字段
	if err := writer.WriteField("master_username", req.MasterUsername); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "写入用户名失败"})
		return
	}
	for _, fileHeader := range req.SessionFiles {
		file, err := fileHeader.Open()
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "打开文件失败"})
			return
		}
		defer file.Close()

		// 创建文件表单字段
		part, err := writer.CreateFormFile("session_files[]", fileHeader.Filename)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "创建文件字段失败"})
			return
		}

		// 复制文件内容到请求体
		if _, err := io.Copy(part, file); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "复制文件内容失败"})
			return
		}
	}
	// 完成multipart写入
	if err := writer.Close(); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "关闭multipart写入器失败"})
		return
	}
	// 创建转发请求
	targetURL := "http://localhost:9090/api/import_session"
	resp, err := http.NewRequest("POST", targetURL, body)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "创建请求失败"})
		return
	}
	resp.Header.Set("Content-Type", writer.FormDataContentType())
	// 发送请求
	client := &http.Client{}
	resp1, err := client.Do(resp)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "请求转发失败"})
		return
	}
	defer resp1.Body.Close()
	// 检查响应状态
	if resp1.StatusCode != http.StatusOK {
		c.JSON(resp1.StatusCode, gin.H{"error": "目标服务返回错误"})
		return
	}
	response.Result(0, nil, "添加成功", c)

}
func (m *TelegramMyApi) DeleteTelegramApi(c *gin.Context) {
	var deleteSessionRequest request.DeleteSessionRequest
	if err := c.ShouldBindJSON(&deleteSessionRequest); err != nil {
		response.FailWithMessage(err.Error(), c)
	}
	fmt.Println(deleteSessionRequest)
	req, err := json.Marshal(&deleteSessionRequest)
	if err != nil {
		panic(err)
	}
	resp, err := http.Post("http://localhost:9090/api/delete_session", "application/json", bytes.NewBuffer(req))
	if err != nil {
		response.FailWithMessage(err.Error(), c)

	}
	defer resp.Body.Close()
	response.Result(0, nil, "删除成功", c)
}
func (m *TelegramMyApi) LoginTelegramApi(c *gin.Context) {
	var loginSessionRequest request.LoginSessionRequest
	if err := c.ShouldBindJSON(&loginSessionRequest); err != nil {
		response.FailWithMessage(err.Error(), c)
	}
	req, err := json.Marshal(&loginSessionRequest)
	if err != nil {
		panic(err)
	}
	resp, err := http.Post("http://localhost:9090/api/sign_in", "application/json", bytes.NewBuffer(req))
	if err != nil {
		response.FailWithMessage(err.Error(), c)

	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	fmt.Println(string(body))
	var result response.SignInResponse
	if err := json.Unmarshal(body, &result); err != nil {
		panic(err)
	}
	fmt.Println(result)

	response.Result(0, result.DataTwo.Session, "加载成功", c)

}
func (m *TelegramMyApi) OutLoginTelegramApi(c *gin.Context) {
	var outLoginSessionRequest request.OutLoginSessionRequest

	if err := c.ShouldBindJSON(&outLoginSessionRequest); err != nil {
		response.FailWithMessage(err.Error(), c)
	}
	fmt.Println("登出:", outLoginSessionRequest)
	req, err := json.Marshal(&outLoginSessionRequest)
	if err != nil {
		panic(err)
	}
	resp, err := http.Post("http://localhost:9090/api/sign_out", "application/json", bytes.NewBuffer(req))
	if err != nil {
		response.FailWithMessage(err.Error(), c)

	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	fmt.Println("登出返回值:", string(body))
	var result response.SignInResponse
	if err := json.Unmarshal(body, &result); err != nil {
		panic(err)
	}
	fmt.Println(result)
	response.Result(0, result.DataTwo.Session, "加载成功", c)

}

func (m *TelegramMyApi) ChangeTagTelegramApi(c *gin.Context) {

	var changeTagTelegramRequest request.ChangeTagTelegramRequest
	if err := c.ShouldBindJSON(&changeTagTelegramRequest); err != nil {
		response.FailWithMessage(err.Error(), c)
	}
	req, err := json.Marshal(&changeTagTelegramRequest)

	if err != nil {
		panic(err)
	}
	resp, err := http.Post("http://localhost:9090/api/change_tag", "application/json", bytes.NewBuffer(req))
	if err != nil {
		response.FailWithMessage(err.Error(), c)

	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	var result response.ChangeTagTelegramResponse
	if err := json.Unmarshal(body, &result); err != nil {
		panic(err)
	}
	fmt.Println("批量分组返回值：", string(body))
	response.Result(result.Status, result.Data, result.Message, c)
}
