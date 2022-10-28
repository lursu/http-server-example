package data

import (
	"database/sql"
	"http-server-example/config"
	"log"
)

type client struct {
	db *sql.DB
}

type DataLayer interface {
	Users
	Healthy() error
}

func New(cfg config.DataBase) (DataLayer, error) {
	log.Println(cfg.Url())
	db, err := sql.Open("postgres", cfg.Url())
	if err != nil {
		return nil, err
	}

	return &client{
		db: db,
	}, db.Ping()
}

// Healthy used to determine healthcheck of service
func (c *client) Healthy() error {
	return c.db.Ping()
}
