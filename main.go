package main

import (
	"GoChallenge/app"
	"log"
	"os"
)

func main() {
	log.Println("Hello. I'm here")

	conf, err := app.Init()
	if err != nil {
		log.Println("Error while parsing config flags: ", err.Error())
		os.Exit(1)
	}

	a := app.NewApp()

	<-a.Start(conf.Port, conf.Query)
}
