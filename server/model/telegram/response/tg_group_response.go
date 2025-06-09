package response

type TelegramGroupResponse struct {
	Id        int    `json:"id"`
	GroupName string `json:"groupName"`
	UserName  string `json:"userName"`
}
type ListTelegramGroupResponse struct {
	List  []TelegramGroupResponse `json:"list"`
	Total int64                   `json:"total"`
}
