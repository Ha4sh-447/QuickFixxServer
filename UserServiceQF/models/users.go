package models

type Users struct {
	Id       int    `json:"id"`
	Name     string `json:"name"`
	Contact  string `json:"contact"`
	Role     string `json:"role"`
	Email    string `json:"email"`
	Password string `json:"password"`
	Image    string `json:"img"`
}

type UserOrderProd struct {
	UserId      int    `json:"userid"`
	ServiceId   int    `json:"serviceid"`
	Field       string `json:"field"`
	OrderId     string `json:"orderid"`
	DateOrdered string `json:"dateOrdered"`
	Status      int    `json:"status"`
}

type KafkaMsg struct {
	UserOrderProd
	ServiceProd
}
