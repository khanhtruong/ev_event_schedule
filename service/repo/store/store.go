package store

import (
	"event_schedule/config"
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type storeImpl struct {
	DB *gorm.DB
}

func NewStore(cfg *config.Config) Store {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable", cfg.PostgresHost, cfg.PostgresUser, cfg.PostgresPass, cfg.PostgresDB, cfg.PostgresPort)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	return &storeImpl{
		DB: db,
	}
}
