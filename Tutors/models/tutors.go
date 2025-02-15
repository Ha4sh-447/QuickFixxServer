package models

type Tutors struct {
	Id       int    `json:"id"`
	Name     string `json:"name"`
	Contact  string `json:"contact"`
	Role     string `json:"role"`
	Email    string `json:"email"`
	Password string `json:"password"`
	Image    string `json:"image"`
	Subject  string `json:"subject"`
	Rating   int    `json:"rating"`
	Fees     int    `json:"fees"`
}

type TutrosDto struct {
	Name     string `json:"name" validate:"required"`
	Contact  string `json:"contact" validate:"required, gt=0 lt=11"`
	Role     string `json:"role"`
	Email    string `json:"email" validate:"required"`
	Password string `json:"password" validate:"required"`
	Image    string `json:"image"`
	Subject  string `json:"subject" validate:"required"`
	Rating   string `json:"rating" validate:"required"`
	Fees     string `json:"rating" validate:"required"`
}
