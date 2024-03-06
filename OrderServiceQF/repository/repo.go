package repository

import (
	"OrderServiceQF/dtos"
	"OrderServiceQF/models"
	"log"

	"github.com/jmoiron/sqlx"
)

type Repo struct {
	*sqlx.DB
}

func NewRepo(db *sqlx.DB) *Repo {
	return &Repo{
		db,
	}
}

func (r *Repo) GetAllOrders() (*[]dtos.OrdersDto, error) {

	var orders []models.Orders
	var ordersDto []dtos.OrdersDto
	err := r.Select(&orders, "SELECT id, orderid, userid, serviceid, field, dateordered, status FROM orders")
	log.Println("HELLO")

	if err != nil {
		log.Println(err)
		return nil, err
	}

	// Initialize ordersDto slice with the same length as orders slice
	ordersDto = make([]dtos.OrdersDto, len(orders))

	for i, order := range orders {
		log.Println("ORDERS", order.OrderId)
		ordersDto[i] = dtos.OrderToDto(&order)
	}

	return &ordersDto, nil
}

func (r *Repo) CreateOrder(odto *dtos.OrdersDto) (string, error) {
	tx := r.MustBegin()
	_, err := tx.Exec("INSERT INTO ORDERS(ORDERID, USERID, SERVICEID, FIELD, DATEORDERED, STATUS) VALUES(?, ?, ? , ? , ?, ?)", odto.OrderId, odto.UserId, odto.ServiceId, odto.Field, odto.DateOrdered, odto.Status)
	log.Println("CREATE ORDER", odto.OrderId)

	if err != nil {
		tx.Rollback()
		return "", err
	}

	err = tx.Commit()
	if err != nil {
		return "", err
	}
	return odto.OrderId, nil
}

func (r *Repo) GetOrderByUserId(userid string) (*[]dtos.OrdersDto, error) {
	tx := r.MustBegin()
	var ordersDto []dtos.OrdersDto
	var orders []models.Orders
	err := tx.Select(&orders, "SELECT id, orderid, userid, serviceid, field, dateordered, status FROM orders WHERE userid=?", userid)

	if err != nil {
		log.Println(err)
		return nil, err
	}

	ordersDto = make([]dtos.OrdersDto, len(orders))

	for i, order := range orders {
		log.Println("ORDERS", order.OrderId)
		ordersDto[i] = dtos.OrderToDto(&order)
	}

	return &ordersDto, nil
}

func (r *Repo) CancelOrder(orderid string) error {
	tx := r.MustBegin()
	_, err := tx.Exec("UPDATE ORDERS SET status=? WHERE orderid=?", -1, orderid)

	if err != nil {
		log.Println("ERROR CANCELLING ORDER", err)
		return err
	}

	if err := tx.Commit(); err != nil {
		log.Println("ERROR COMMITTING TRANSACTION:", err)
		return err
	}

	return nil
}
