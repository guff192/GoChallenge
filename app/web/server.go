package web

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"time"
)

type Server struct {
	done chan bool
}

func (s Server) Start(port string) {
	s.setUpHandlers()

	go func() {
		Next = time.After(10 * time.Second)
		if err := http.ListenAndServe(":"+port, nil); err != nil {
			log.Println("Failed to start web: ", err)
		}
	}()
}

func (s Server) setUpHandlers() {
	http.HandleFunc("/", DefaultHandler)
}

func DefaultHandler(w http.ResponseWriter, r *http.Request) {
	select {
	case GithubResp = <-GithubResponseChannel:
		ParseFinished = true
	case <-Next:
		TimeElapsed = true
	default:

	}
	if TimeElapsed && ParseFinished {
		t, _ := template.ParseFiles("index.html")
		t.Execute(w, GithubResp)
	} else {
		fmt.Fprintf(w, "Идёт поиск")
	}
}
