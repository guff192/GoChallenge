package app

import (
	"fmt"
	"log"
	"net/http"
)

type Server struct {
	done chan bool
}

func (s Server) Start() {
	s.setUpHandlers()

	go func() {
		if err := http.ListenAndServe(":8888", nil); err != nil {
			log.Println("Failed to start server: ", err)
		}
	}()
}

func (s Server) setUpHandlers() {
	http.HandleFunc("/", DefaultHandler)
}

func DefaultHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, world!")
}
