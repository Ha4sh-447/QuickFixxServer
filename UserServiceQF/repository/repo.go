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
	log.Println("UPDATE_USER_REPO", user)
	_, err := tx.Exec("UPDATE T_USERSDB SET name=?, contact=?, email=?, password=? WHERE id=?",
		user.Name, user.Contact, user.Email, user.Password, user.Id)
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

func (r *Repo) GetUserByEmail(email string) (*models.Users, error) {
	tx := r.repo.MustBegin()
	var user models.Users

	err := tx.Get(&user, "SELECT * FROM t_usersdb WHERE email=?", email)
	if err != nil {
		log.Println("ERROR", err)
		return nil, err
	}

	return &user, nil
}

func (r *Repo) GetSpDetails(spId int32, field string) (*models.ServiceProd, error) {

	var table string

	switch field {
	case "electrician":
		// port = "8080"
		table = "t_electriciandb_2"
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
	query := fmt.Sprintf("SELECT * FROM %s WHERE eid=?", table)

	err := tx.Get(&sp, query, spId)
	log.Println("SP DETAILS", sp)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	return &sp, nil
}

func (r *Repo) GetUserContactById(id int32) (string, string, error) {
	tx := r.repo.MustBegin()

	result := tx.QueryRow("SELECT name,contact FROM t_usersdb WHERE id=?", id)
	var contact string
	var name string
	err := result.Scan(&name, &contact)
	log.Println(contact + " " + name)
	if err != nil {
		log.Println(err)
		return "", "", err
	}

	return name, contact, nil
}
