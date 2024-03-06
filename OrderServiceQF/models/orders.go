package models

type Orders struct {
	// Id          int    `json:"id"`
	// OrderId     string `json:"orderId"`
	// UserId      int    `json:"userId"`
	// ServiceId   int    `json:"serviceId"`
	// DateOrdered string `json:"dateOrdered"`
	// Field       string `json:"field"`
	// Status      bool   `json:"status"`
	Id          int    `json:"id"`
	OrderId     string `json:"orderId"`
	UserId      int64  `json:"userId"`
	ServiceId   string `json:"serviceId"`
	Field       string `json:"field"`
	DateOrdered string `json:"dateOrdered"`
	Status      int    `json:"status"`
}

type OrderResponse struct {
	UserName   string `json:"user_name"`
	ProName    string `json:"pro_name"`
	ProContact string `json:"pro_contact"`
	ProServ    string `json:"pro_service"`
	ProRating  string `json:"pro_rating"`
}

type UserOrderProd struct {
	UserId      int    `json:"userid"`
	ServiceId   int    `json:"serviceid"`
	Field       string `json:"field"`
	OrderId     string `json:"orderid"`
	DateOrdered string `json:"dateOrdered"`
	Status      int    `json:"status"`
}
