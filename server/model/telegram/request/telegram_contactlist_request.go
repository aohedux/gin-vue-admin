package request

type ListTelegramContactlistRequest struct {
	Page     int    `json:"page"`
	PageSize int    `json:"page_size"`
	Phone    string `json:"phone"`
}
