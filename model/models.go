package model

type UserBalance struct {
	Id      int  `json:"-"`
	Balance uint `json:"balance"`
}

type Reservation struct {
	UserId  int  `json:"id"`
	TaskId  int  `json:"task_id"`
	OrderId int  `json:"order_id"`
	Amount  uint `json:"balance"`
}
