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

type ServiceProd struct {
	Id            int     `json:"id"`
	Address       string  `json:"address"`
	Contact       string  `json:"contact"`
	Name          string  `json:"name"`
	Experience    string  `json:"experience"`
	Qualification []uint8 `json:"qualification"`
	// Qualification  interface{} `json:"qualification"`
	Location       string `json:"location"`
	Rating         int    `json:"rating"`
	Specialization string `json:"specialization"`
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
	Username string `json:"username"`
}
