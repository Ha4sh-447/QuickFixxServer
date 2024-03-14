package main

import (
	"UserServiceQF/controller"
	"UserServiceQF/database"
	"UserServiceQF/kafka"
	"UserServiceQF/logs"
	"UserServiceQF/repository"
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {

	fmt.Println("SERVER")

	//l := log.New(os.Stdout, "USER-SERVICE", log.LstdFlags)

	//	Load env file
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("SERVER")

	//Logger config
	logrConfig := logs.InitLogger(
		os.Getenv("LOG_FILE_NAME"),
		10,
		10,
		100,
	)
	logrConfig.LoadLogger()

	config := &database.Config{
		Host:     os.Getenv("DB_HOST"),
		Port:     os.Getenv("DB_PORT"),
		Password: os.Getenv("DB_PASS"),
		User:     os.Getenv("DB_USER"),
		DBName:   os.Getenv("DB_NAME"),
	}
	fmt.Println("SERVER")

	producer, err := kafka.SetupProducer()
	if err != nil {
		log.Println(err)
	}

	newSqlDb, err := config.ConnectToDb()
	if err != nil {
		log.Panic(err)
	}

	repo := repository.RepoInit(newSqlDb)
	uh := controller.NewController(repo, &producer)

	fmt.Println("SERVER")
	var port string
	if p := os.Getenv("SERVER_PORT"); p != "" {
		port = fmt.Sprintf(":%s", p)
	} else {
		port = ":9000"
	}

	// order := models.UserOrderProd{
	// 	UserId:      7,
	// 	ServiceId:   1,
	// 	Field:       "Electrician",
	// 	OrderId:     "order_1111",
	// 	DateOrdered: "02/03/2024",
	// 	Status:      0,
	// }

	r := gin.Default()
	user := r.Group("/api/users")
	user.GET("/", uh.GetAllUsers)
	user.GET("/email", uh.GetByEmail)
	user.POST("/", uh.CreateUser)
	user.POST(("/send"), uh.PostOrder)
	user.DELETE("/:id", uh.DeleteUser)
	user.GET("/:id", uh.GetUserById)
	user.PUT("/:id", uh.UpdateUser)

	srv := &http.Server{
		Addr:    port,
		Handler: r,
	}
	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen: %s\n", err)
		}
	}()

	// Graceful shutdown
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	log.Println("Shutting down server...")

	// Create a context with a timeout
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// Attempt to gracefully close connections
	if err := srv.Shutdown(ctx); err != nil {
		log.Fatalf("Server forced to shutdown: %v\n", err)
	}
	log.Println("Server exiting")
}

//s := http.Server{
//	Addr:                         ":9090",
//	Handler:                      r,
//	DisableGeneralOptionsHandler: false,
//	TLSConfig:                    nil,
//	ReadTimeout:                  5 * time.Second,
//	ReadHeaderTimeout:            0,
//	WriteTimeout:                 10 * time.Second,
//	IdleTimeout:                  120 * time.Second,
//	MaxHeaderBytes:               0,
//	TLSNextProto:                 nil,
//	ConnState:                    nil,
//	ErrorLog:                     l,
//	BaseContext:                  nil,
//	ConnContext:                  nil,
//}
//
////	to implement
////listening to server
////gracefully shutting down
//
////listenAndServe going to block the code from proceeding further, so put it in a coroutine
//go func() {
//	l.Println("Starting server on port 9090.")
//	err := s.ListenAndServe()
//	if err != nil {
//		l.Printf("Error listening server: %s\n", err)
//		os.Exit(1)
//	}
//}()
//
////	Graceful shutdown
////	trap the signature and then shutdown server
//c := make(chan os.Signal, 1)
//signal.Notify(c, os.Interrupt)
//signal.Notify(c, os.Kill)
//
////	Block this part until the signal is recieved, only then proceed further
//sig := <-c
//l.Println("Got signal: ", sig)
//l.Println("SERVER CLOSED")
//
//// gracefully shutdown the server, waiting max 30 seconds for current operations to complete
//ctx, _ := context.WithTimeout(context.Background(), 30*time.Second)
//s.Shutdown(ctx)
