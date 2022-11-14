package model

type User struct {
	ID    string `json:"id"`
	Name  string `json:"name"`
	Age   uint32 `json:"age"`
	Email string `json:"email"`
}
