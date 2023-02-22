package handler

import (
	"event_schedule/config"
	"event_schedule/service/repo"
)

type Handler struct {
	Config *config.Config
	Repo   repo.Repo
}

func NewHandler(cfg *config.Config) Handler {
	r := repo.NewRepo(cfg)

	return Handler{
		Config: cfg,
		Repo:   r,
	}
}
