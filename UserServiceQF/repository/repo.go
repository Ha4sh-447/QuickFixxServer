package repository

import (
	"UserServiceQF/dto"
	"UserServiceQF/models"
	"fmt"
	"log"

	"github.com/jmoiron/sqlx"
	logr "github.com/sirupsen/logrus"
)

type Repo struct {
	repo *sqlx.DB
}

func RepoInit(repos *sqlx.DB) *Repo {
	return &Repo{
		repos,
	}
}

func (r *Repo) GetAll() ([]models.Users, error) {
	var users []models.Users

	err := r.repo.Select(&users, "SELECT * FROM T_USERSDB")
	if err != nil {
		log.Panic(err)
	}

	return users, nil
}

func (r *Repo) GetById(id string) (*dto.UserDto, error) {
	tx := r.repo.MustBegin()
	row := tx.QueryRowx("select * from t_usersdb where id=$1", id)
	var usrDto dto.UserDto
	err := row.Scan(&usrDto.Name, &usrDto.Contact, &usrDto.Email, &usrDto.Image)
	if err != nil {
		return nil, err
	}
	return &usrDto, nil
}

func (r *Repo) AddUsr(usr *dto.UserDto) error {
	tx := r.repo.MustBegin()
	_, err := tx.Exec("INSERT INTO T_USERSDB(name, role, contact, email, password, image) VALUES (?, ?, ?, ?, ?, ?)", usr.Name, usr.Role, usr.Contact, usr.Email, usr.Password, usr.Image)
	log.Println("REPO", "ADD USER")
	logr.Debugf("Added user: Name=%s, Role=%s, Contact=%s, Email=%s, Image=%s\n", usr.Name, usr.Role, usr.Contact, usr.Email, usr.Image)
	if err != nil {
		tx.Rollback()
		return err
	}

	err = tx.Commit()
	if err != nil {
		return err
	}

	return nil
}

func (r *Repo) DeleteUsr(id string) error {
	tx := r.repo.MustBegin()

	_, err := tx.Exec("DELETE FROM t_usersdb WHERE id = ?", id)
	if err != nil {
		tx.Rollback()
		return err
	}
	logr.Debugf("Delete user %s", id)

	err = tx.Commit()
	if err != nil {
		return err
	}

	return nil
}

func (r *Repo) UpdateUsr(user *models.Users) error {
	// Start a new transaction
	tx := r.repo.MustBegin()

	_, err := tx.Exec("UPDATE T_USERSDB SET name=$1, contact=$2, email=$3, password=$4, img=$5 WHERE id=$6",
		user.Name, user.Contact, user.Email, user.Password, user.Image, user.Id)
	if err != nil {
		tx.Rollback()
		return err
	}

	err = tx.Commit()
	if err != nil {
		return err
	}

	return nil
}

func (r *Repo) GetSpDetails(spId int32, field string) (*models.ServiceProd, error) {

	var table string

	switch field {
	case "electrician":
		// port = "8080"
		table = "t_electriciandb"
	case "plumber":
		// port = "8083"
		table = "t_plumbers"
	case "carpenter":
		// port = "8082"
		table = "t_carpenter"
	case "housekeeping":
		// port = "joemama"
		table = "t_housekeeping"
	}

	// call := fmt.Sprintf("http://localhost:%s/api/%s/%d", port, field, spId)
	// log.Println(call)

	tx := r.repo.MustBegin()
	defer tx.Rollback()
	log.Println("INPUTS", spId, field)
	var sp models.ServiceProd
	query := fmt.Sprintf("SELECT * FROM %s WHERE id=?", table)

	err := tx.Get(&sp, query, spId)
	log.Println("SP DETAILS", sp)
	if err != nil {
		return nil, err
	}

	return &sp, nil
}
