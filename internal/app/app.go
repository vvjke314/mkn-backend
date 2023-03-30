package app

import "context"

type Application struct {
	ctx context.Context
	//repo repository.Repository #ADD REPO via GORM
}

func New(ctx context.Context) *Application {
	return &Application{
		ctx: ctx,
	}
}

func (a *Application) Run() {
	a.StartServer()
}
