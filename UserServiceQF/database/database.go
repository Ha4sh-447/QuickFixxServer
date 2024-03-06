package database

import (
	"context"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"time"
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

func (c *Config) ConnectToDb() (*sqlx.DB, error) {

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	conn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", c.User, c.Password, c.Host, c.Port, c.DBName)
	db, err := sqlx.Connect("mysql", conn)

	if err != nil {
		return nil, err
	}

	if err := db.PingContext(ctx); err != nil {
		db.Close()
		return nil, err
	}

	return db, nil
}
