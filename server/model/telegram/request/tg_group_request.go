package request

type TelegramGroupRequest struct {
	GroupName string `json:"groupName"`
	UserName  string `json:"userName"`
}

type ListTelegramGroupRequest struct {
	Page      int    `json:"page"`
	PageSize  int    `json:"page_size"`
	UserName  string `json:"user_name"`
	GroupName string `json:"group_name"`
}

type UpdateTelegramGroupRequest struct {
	Id        int    `json:"id"`
	GroupName string `json:"groupName"`
	UserName  string `json:"userName"`
}
type DeleteTelegramGroupRequest struct {
	Id []int `json:"ids"`
}
