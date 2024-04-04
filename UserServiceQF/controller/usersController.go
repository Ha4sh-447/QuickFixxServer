package controller

import (
	"UserServiceQF/dto"
	"UserServiceQF/kafka"
	"UserServiceQF/models"
	"UserServiceQF/repository"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/IBM/sarama"
	"github.com/gin-gonic/gin"
	logr "github.com/sirupsen/logrus"
)

type UserCont struct {
	*repository.Repo
	producer *sarama.SyncProducer
}

func NewController(repo *repository.Repo, producer *sarama.SyncProducer) *UserCont {
	return &UserCont{
		repo,
		producer,
	}
}

type Product struct {
	ID          int     `json:"id" `
	Name        string  `json:"name" validate:"required"`
	Description string  `json:"desc"`
	Price       float32 `json:"price" validate:"gt=0"`
	SKU         string  `json:"sku" validate:"required, sku"`
	CreatedOn   string  `json:"-"`
	UpdatedOn   string  `json:"-"`
	DeletedOn   string  `json:"-"`
}

type Products []*Product

func (u *UserCont) GetAllUsers(ctx *gin.Context) {

	users, err := u.Repo.GetAll()

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, dto.Error{Code: -1, Message: err.Error()})
		logr.Debug("[ERROR]: ", err)
	} else {
		ctx.JSON(http.StatusOK, users)
	}

	resp, err := http.Get("http://localhost:9000/")
	if err := json.NewDecoder(resp.Body).Decode(&Products{}); err != nil {
		return
	}
}

func (u *UserCont) GetUserById(ctx *gin.Context) {
	user, err := u.Repo.GetById(ctx.Params.ByName("id"))

	if err == nil {
		ctx.JSON(http.StatusOK, user)
	} else {
		ctx.JSON(http.StatusBadRequest, dto.Error{
			Code:    -1,
			Message: err.Error(),
		})
	}
}

func (u *UserCont) DeleteUser(ctx *gin.Context) {
	id := ctx.Params.ByName("id")
	err := u.DeleteUsr(id)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, dto.Error{-1, err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, dto.Response{Message: fmt.Sprintf("User %s deleted", id)})
}

func (u *UserCont) CreateUser(ctx *gin.Context) {
	var usrdto dto.UserDto

	if err := ctx.ShouldBindJSON(&usrdto); err != nil {
		ctx.JSON(http.StatusBadRequest, dto.Error{Code: -1, Message: "Invalid request body"})
		return
	}

	log.Println("USER-CONT", usrdto)

	if err := u.AddUsr(&usrdto); err != nil {
		ctx.JSON(http.StatusInternalServerError, dto.Error{Code: -1, Message: err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, dto.Response{Message: "User created successfully"})
}

func (u *UserCont) UpdateUser(ctx *gin.Context) {
	id := ctx.Params.ByName("id")

	var usrdto dto.UserDto
	if err := ctx.ShouldBindJSON(&usrdto); err != nil {
		ctx.JSON(http.StatusBadRequest, dto.Error{Code: -1, Message: "Invalid request body"})
		return
	}
	uid, _ := strconv.Atoi(id)
	fmt.Println(uid)
	user := models.Users{
		Id:       uid,
		Name:     usrdto.Name,
		Contact:  usrdto.Contact,
		Email:    usrdto.Email,
		Password: usrdto.Password,
		Image:    usrdto.Image,
	}

	// Update the user in the database
	if err := u.UpdateUsr(&user); err != nil {
		log.Println("UPDATE_USER_ERROR", err)
		ctx.JSON(http.StatusInternalServerError, dto.Error{Code: -1, Message: "Failed to update user"})
		return
	}

	ctx.JSON(http.StatusOK, dto.Response{Message: "User updated successfully"})
}

func (u *UserCont) GetByEmail(ctx *gin.Context) {
	email := ctx.Query("email")

	user, err := u.Repo.GetUserByEmail(email)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, dto.Error{Code: -1, Message: err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, user)
}

func (u *UserCont) PostOrder(ctx *gin.Context) {
	var userorder models.UserOrderProd
	if err := ctx.ShouldBindJSON(&userorder); err != nil {
		log.Println(err)
		ctx.JSON(http.StatusBadRequest, dto.Error{Code: -1, Message: "Invalid request body"})
		return
	}

	// Fetch data in parallel
	type Result struct {
		SP       *models.ServiceProd
		Username string
		Contact  string
		Err      error
	}
	resultCh := make(chan Result, 2) // Buffer for 2 results

	go func() {
		sp, err := u.GetSpDetails(int32(userorder.ServiceId), userorder.Field)
		resultCh <- Result{SP: sp, Err: err}
	}()

	go func() {
		username, contact, err := u.Repo.GetUserContactById(int32(userorder.UserId))
		resultCh <- Result{Username: username, Contact: contact, Err: err}
	}()

	// Wait for both results
	var sp *models.ServiceProd
	var username string
	var contact string
	for i := 0; i < 2; i++ {
		result := <-resultCh
		if result.Err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch data"})
			return
		}
		if result.SP != nil {
			sp = result.SP
		}
		if result.Username != "" {
			username = result.Username
		}
		if result.Contact != "" {
			contact = result.Contact
		}
	}

	kafkaOrder := models.KafkaMsg{
		UserOrderProd: userorder,
		ServiceProd:   *sp,
		Username:      username,
		Contact:       contact,
	}

	err := kafka.SendKafkaMessage(*u.producer, kafkaOrder)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to send Kafka message"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "Message sent"})
}
