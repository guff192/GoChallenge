package app

import (
	"GoChallenge/app/web"
)

type App struct {
	done   <-chan bool
	server web.Server
	client web.Client
}

func NewApp() *App {
	return &App{
		done: make(chan bool),
	}
}

func (a *App) Start(port, query string) <-chan bool {
	a.server.Start(port)
	a.client.RunQuery(query)

	return a.done
}
