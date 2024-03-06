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

//type DB struct {
//	*sqlx.DB
//}

func (c *Config) ConnectDB() (*sqlx.DB, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	conn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", c.User, c.Password, c.Host, c.Port, c.DBName)
	fmt.Println("Connecting to database with connection string:", conn) // Log connection string
	db, err := sqlx.Connect("mysql", conn)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to database: %v", err)
	}

	if err := db.PingContext(ctx); err != nil {
		db.Close()
		return nil, fmt.Errorf("failed to ping database: %v", err)
	}

	fmt.Println("Successfully connected to database")
	return db, nil
}
