package response

import "time"

type PageResult struct {
	List     interface{} `json:"list"`
	Total    int64       `json:"total"`
	Page     int         `json:"page"`
	PageSize int         `json:"pageSize"`
}

// 响应时查询所有
type Session struct {
	Phone     string    `json:"phone"`
	Username  string    `json:"username"`
	Name      string    `json:"name"`
	IsPremium bool      `json:"is_premium"`
	IsBlocked bool      `json:"is_blocked"`
	Online    bool      `json:"online"`
	Server    int       `json:"server"`
	Proxy     string    `json:"proxy"`
	CreatedAt time.Time `json:"created_at"`
	Tag       int       `json:"tag"`
}

type Data struct {
	Session []Session `json:"session"`
	Count   int       `json:"count"`
}

type GetSessionResponse struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
	Data    Data   `json:"data"`
}
type SignInResponse struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
	DataTwo Data   `json:"data"`
}
type DataTwo struct {
	Session []Session `json:"session"`
}
type GetSessionResponseTwo struct {
	Session2 []Session2 `json:"session"`
	Count    int        `json:"count"`
}
type Session2 struct {
	Phone     string    `json:"phone"`
	Username  string    `json:"username"`
	Name      string    `json:"name"`
	IsPremium bool      `json:"is_premium"`
	IsBlocked bool      `json:"is_blocked"`
	Online    bool      `json:"online"`
	Server    int       `json:"server"`
	Proxy     string    `json:"proxy"`
	CreatedAt time.Time `json:"created_at"`
	Tag       string    `json:"tag"`
}

type ChangeTagTelegramResponse struct {
	Status  int                   `json:"status"`
	Message string                `json:"message"`
	Data    ChangeTagTelegramData `json:"data"`
}

//批量分组数据
type ChangeTagTelegramData struct {
	FailureCount int32 `json:"failure_count"`
	SuccessCount int32 `json:"success_count"`
}
