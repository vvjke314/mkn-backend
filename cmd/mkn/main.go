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
func main() {
	ctx := context.Background()
	a := app.New(ctx)
	a.Run()
}
