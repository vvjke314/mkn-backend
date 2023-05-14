package main

import (
	"context"
	"mkn-backend/internal/app"
)

// @title          	MKN API
// @version         1.0
// @description     Notification backend service.
// @contact.name   MKN Support
// @contact.email  mkn-notifyer@mail.ru
// @host      127.0.0.1:8080
// @BasePath  /
// @securityDefinitions.apikey BearerAuth
// @in header
// @name Authorization
// @schemes http
func main() {
	ctx := context.Background()
	a, err := app.New(ctx)
	if err != nil {
		panic("Can't create application")
	}
	a.Run()
}
