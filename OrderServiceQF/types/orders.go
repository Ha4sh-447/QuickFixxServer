package types

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
	Address        string  `json:"address"`
	Contact        string  `json:"contact"`
	Specialization string  `json:"specialization"`
	Eid            int     `json:"eid"`
	Specz          string  `json:"specz"`
	Rating         float32 `json:"rating"`
	Experience     string  `json:"experience"`
	U_id           int     `json:"u_id"`
	Shopname       string  `json:"shopname"`
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
	Contact  string `json:"contact"`
}
