package web

import (
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
		if err := http.ListenAndServe(":"+port, nil); err != nil {
			log.Println("Failed to start web: ", err)
		}
	}()
}

func (s Server) setUpHandlers() {
	http.HandleFunc("/", DefaultHandler)
}

func DefaultHandler(w http.ResponseWriter, r *http.Request) {
	if delayer == nil { // create delayer if it doesn't exist
		delayer = time.NewTicker(10 * time.Second)
	}

	select {
	case <-delayer.C:
		// 10 seconds elapsed
		timeElapsed = true
		delayer.Stop()
	case githubResp = <-GithubResponseChannel:
		// client received response from github
	default:
	}

	if timeElapsed && len(githubResp.Items) > 0 {
		t, _ := template.ParseFiles("index.html")
		w.WriteHeader(http.StatusOK)
		t.Execute(w, githubResp)
	} else {
		w.Write([]byte("Идёт поиск"))
	}
}
