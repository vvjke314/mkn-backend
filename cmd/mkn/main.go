package main

import (
	"context"
	"log"
	"mkn-backend/internal/app"
	"mkn-backend/internal/pkg/grpcApi"
	"net"
	"sync"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
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
	var wg sync.WaitGroup
	ctx := context.Background()
	a, err := app.New(ctx)
	if err != nil {
		panic("Can't create application")
	}

	server, err := grpcApi.New(ctx)
	if err != nil {
		panic("Can't create grpc server")
	}

	wg.Add(1)
	go func(a *app.Application) {
		defer wg.Done()
		a.Run()
	}(a)

	wg.Add(1)
	go func(service *grpcApi.GRPCServer) {
		defer wg.Done()
		s := grpc.NewServer()
		reflection.Register(s)
		srv := service
		grpcApi.RegisterBackendServiceServer(s, srv)
		log.Println("GRPC SERVER STARTED")
		//зарегистрировать сервисы
		l, err := net.Listen("tcp", ":9000")
		if err != nil {
			log.Fatal(err)
		}

		if err := s.Serve(l); err != nil {
			log.Fatal(err)
		}

	}(server)

	wg.Wait()
}
