package app

type App struct {
	done   <-chan bool
	server Server
	client Client
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
