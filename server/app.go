package server

import (
	"http-server-example/config"
	"http-server-example/data"
	"time"
)

type app struct {
	start time.Time
	db    data.DataLayer
}

func New(cfg config.Config) (*app, error) {
	db, err := data.New(cfg.DataBase)
	if err != nil {
		return nil, err
	}

	return &app{
		start: time.Now(),
		db:    db,
	}, nil
}

func (a *app) Healthy() error {
	return a.db.Healthy()
}
