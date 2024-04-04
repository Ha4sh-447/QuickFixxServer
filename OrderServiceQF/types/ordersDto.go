package types

type OrdersDto struct {
	OrderId     string `json:"orderid"`
	UserId      int64  `json:"userid"`
	DateOrdered string `json:"dateOrdered"`
	ServiceId   string `json:"serviceid"`
	Field       string `json:"field"`
	Status      int    `json:"status"`
}

func OrderToDto(orders *Orders) OrdersDto {
	return OrdersDto{
		OrderId:     orders.OrderId,
		UserId:      orders.UserId,
		DateOrdered: orders.DateOrdered,
		ServiceId:   orders.ServiceId,
		Field:       orders.Field,
		Status:      orders.Status,
	}
}
