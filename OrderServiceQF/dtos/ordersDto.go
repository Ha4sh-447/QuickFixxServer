package dtos

import (
	"OrderServiceQF/models"
)

type OrdersDto struct {
	//ID        string    `json:"id"`
	OrderId string `json:"orderid"`
	UserId  int64  `json:"userid"`
	//ServiceId string    `json:"service_id"`
	DateOrdered string `json:"dateOrdered"`
	ServiceId   string `json:"serviceid"`
	Field       string `json:"field"`
	Status      int    `json:"status"`
	// ID          string    `json:"id"`
	// OrderId     string    `json:"orderid"`
	// UserId      string    `json:"userid"`
	// ServiceId   string    `json:"serviceid"`
	// DateOrdered time.Time `json:"dateOrdered"`
	// Field       string    `json:"field"`
	// Status      bool      `json:"status"`
}

func OrderToDto(orders *models.Orders) OrdersDto {
	return OrdersDto{
		OrderId:     orders.OrderId,
		UserId:      orders.UserId,
		DateOrdered: orders.DateOrdered,
		ServiceId:   orders.ServiceId,
		Field:       orders.Field,
		Status:      orders.Status,
	}
}
