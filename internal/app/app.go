package app

import (
	"context"
	"fmt"
	"mkn-backend/internal/pkg/config"
	"mkn-backend/internal/pkg/redisClient"
	"mkn-backend/internal/pkg/repository"

	"github.com/redis/go-redis/v9"
	log "github.com/sirupsen/logrus"
)

type Application struct {
	ctx    *context.Context
	repo   *repository.Repository
	redis  *redis.Client
	config *config.Config
}

func New(ctx context.Context) (*Application, error) {
	cfg, err := config.NewConfig(ctx)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable", cfg.DataBase.Host, cfg.DataBase.User, cfg.DataBase.Password, cfg.DataBase.Name, cfg.DataBase.Port)
	repo, err := repository.New(dsn)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	redis := redisClient.New()

	return &Application{
		ctx:    &ctx,
		repo:   repo,
		redis:  redis,
		config: cfg,
	}, nil
}

func (a *Application) Run() {
	log.Println("Application started")
	a.StartServer()
	log.Println("Application shutted down")
}
