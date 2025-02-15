package database

import (
	"context"
	"fmt"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	logr "github.com/sirupsen/logrus"
)

type Config struct {
	Host     string
	Port     string
	Password string
	User     string
	DBName   string
}

type DB struct {
	Connection *sqlx.DB
}

func NewDB(sqlConn *sqlx.DB) *DB {
	return &DB{
		sqlConn,
	}
}

func CreateTable(db *sqlx.DB) error {
	schema := `CREATE TABLE IF NOT EXISTS t_tutors(
	 	id INT AUTO_INCREMENT PRIMARY KEY,
		name VARCHAR(100) NOT NULL,
		contact VARCHAR(50) NOT NULL,
		role VARCHAR(50) NOT NULL,
		email VARCHAR(100) UNIQUE NOT NULL,
		password TEXT NOT NULL,
		image TEXT
		subject VARCHAR(30) NOT NULL,
		rating INT,
		fees INT NOT NULL
	)`

	_, err := db.Exec(schema)
	if err != nil {
		logr.Errorf("Error creating tutors table %v", err)
		return err
	}

	return nil
}

func (c *Config) ConnectToDb() (*sqlx.DB, error) {

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	fmt.Println("Connecting to:", fmt.Sprintf("%s:%stcp(%s:%s)/%s", c.User, c.Password, c.Host, c.Port, c.DBName))

	conn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", c.User, c.Password, c.Host, c.Port, c.DBName)
	db, err := sqlx.Connect("mysql", conn)

	if err != nil {
		return nil, err
	}
	CreateTable(db)

	if err := db.PingContext(ctx); err != nil {
		db.Close()
		return nil, err
	}

	return db, nil
}
