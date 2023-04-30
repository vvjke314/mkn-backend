package main

import (
	"context"
	"fmt"
	"mkn-backend/internal/pkg/config"

	"github.com/BurntSushi/toml"
	_ "github.com/lib/pq"
	"github.com/pressly/goose"
	log "github.com/sirupsen/logrus"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

const (
	migrationsPath = "migrations"
	driver         = "postgres"
)

func main() {
	ctx := context.Background()

	log.Info("Starting migrations")
	var conf config.Config
	if _, err := toml.DecodeFile("config.toml", &conf); err != nil {
		log.WithContext(ctx).WithError(err).Error("No 'config.toml' file loaded")
	}

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable", conf.DataBase.Host, conf.DataBase.User, conf.DataBase.Password, conf.DataBase.Name, conf.DataBase.Port)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.WithContext(ctx).WithError(err).Error("Can't connect to database")
	}

	dataBase, err := db.DB()
	if err != nil {
		log.WithContext(ctx).WithError(err).Error("Can't get SQL database")
	}
	log.Info("The database connection was established successfully")

	err = goose.Up(dataBase, migrationsPath)
	if err != nil {
		log.WithContext(ctx).WithError(err).Error("Failed to migrate: %v", err)
	}
	log.Info("Database migrations completed")
}
