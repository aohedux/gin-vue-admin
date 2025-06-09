package response

type ListTelegramContactListResponse struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
	Data    Data   `json:"data"`
}
type TelegramContactListResponse struct {
	TelegramID int    `json:"telegram_id"`
	FirstName  string `json:"first_name"`
	LastName   string `json:"last_name"`
	UserName   string `json:"username"`
}
type Data struct {
	Contacts []TelegramContactListResponse `json:"contacts"`
	Count    int                           `json:"count"`
}
