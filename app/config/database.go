package config

import (
	"os"
	"time"

	"github.com/jmoiron/sqlx"
)

type Client struct {
	DB *sqlx.DB
}

// InitDB connection
func InitClient() (*Client, error) {
	// connect to postgres database
	db, err := sqlx.Connect("postgres", os.Getenv("DB_URI"))
	if err != nil {
		time.Sleep(time.Second * 10)
		db, err = sqlx.Connect("postgres", os.Getenv("DB_URI"))
		if err != nil {
			return nil, err
		}
	}

	return &Client{
		DB: db.Unsafe(),
	}, nil
}
