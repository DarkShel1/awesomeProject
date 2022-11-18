package model

type Reservation struct {
	UserId  int    `json:"user_id"`
	TaskId  int    `json:"task_id"`
	OrderId uint32 `json:"order_id"`
	Amount  int    `json:"amount"`
}
