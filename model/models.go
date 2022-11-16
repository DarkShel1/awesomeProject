package model

type User struct {
	ID    string `json:"id"`
	Name  string `json:"name"`
	Age   string `json:"age"`
	Email string `json:"email"`
}

type Balance struct {
	UserID  int32 `json:"user_id"`
	Balance int32 `json:"balance"`
}
