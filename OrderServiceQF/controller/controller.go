package controller

import (
	"OrderServiceQF/repository"
	"OrderServiceQF/types"
	"fmt"
	"log"
	"net/http"

	"github.com/labstack/echo/v4"
)

type Controller struct {
	r *repository.Repo
}

func NewCont(repo *repository.Repo) *Controller {
	return &Controller{
		repo,
	}
}

func (c *Controller) GetAll(e echo.Context) error {
	orders, err := c.r.GetAllOrders()
	if err != nil {
		e.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to get orders"})
		return nil
	}

	e.JSON(http.StatusOK, orders)
	return nil
}

func (c *Controller) PlaceOrder(e echo.Context) error {
	dto := types.OrdersDto{}
	if err := e.Bind(&dto); err != nil {
		log.Printf("\nERROR", err)
		e.JSON(http.StatusPartialContent, "Can't decode json")
		return err
	}
	log.Printf("BODY", dto.OrderId, dto.UserId, dto.ServiceId, dto.Field, dto.DateOrdered, dto.Status)
	result, err := c.r.CreateOrder(&dto)

	if err != nil {
		e.JSON(http.StatusInternalServerError, err)
		return err
	}

	str := fmt.Sprint("Order placed", result)
	e.JSON(http.StatusOK, str)
	return nil
}

func (c *Controller) CancelOrder(e echo.Context) error {
	orderid := e.Param("orderid")
	err := c.r.CancelOrder(orderid)
	if err != nil {
		e.JSON(http.StatusBadRequest, err)
		return err
	}

	log.Println("ORDER CANCELED", orderid)
	e.JSON(http.StatusOK, orderid)
	return nil
}

func (c *Controller) GetOrderByUserId(e echo.Context) error {

	userid := e.Param("userid")
	orders, err := c.r.GetOrderByUserId(userid)
	if err != nil {
		e.JSON(http.StatusInternalServerError, err)
		return err
	}

	e.JSON(http.StatusOK, orders)
	return nil
}
