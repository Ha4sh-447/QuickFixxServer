package main

import (
	"OrderServiceQF/controller"
	"OrderServiceQF/database"
	"OrderServiceQF/kafka"
	"OrderServiceQF/repository"
	models "OrderServiceQF/types"
	"context"
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
)

func main() {

	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal(err)
	}

	config := &database.Config{
		Host:     os.Getenv("DB_HOST"),
		Port:     os.Getenv("DB_PORT"),
		Password: os.Getenv("DB_PASS"),
		User:     os.Getenv("DB_USER"),
		DBName:   os.Getenv("DB_NAME"),
	}

	fmt.Println("---", config.Port)
	conn, err := config.ConnectDB()
	if err != nil {
		log.Panic("DB ERROR", err.Error())
	}

	repo := repository.NewRepo(conn)
	c := controller.NewCont(repo)

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	go kafkaSetup(ctx)

	e := echo.New()
	path := e.Group("/v1/orders")
	path.GET("/", c.GetAll)
	path.GET("/:userid", c.GetOrderByUserId)
	path.DELETE("/:orderid", c.CancelOrder)
	path.POST("/", c.PlaceOrder)
	go func() {
		if err := e.Start(":1323"); err != nil {
			e.Logger.Fatal(err)
		}
	}()

	// Graceful shutdown
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	cancel()
	if err := e.Shutdown(ctx); err != nil {
		e.Logger.Fatal(err)
	}
}

func kafkaSetup(ctx context.Context) {
	order := models.UserOrderProd{} // Provide an initial order if needed
	kafka.SetupConsumerGroup(ctx, order)
}
