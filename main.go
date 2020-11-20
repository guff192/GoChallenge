package main

import (
	"GoChallenge/app"
	"flag"
	"log"
	"os"
)

const (
	defaultPort  = "8888"
	defaultQuery = "golang"
)

func main() {
	log.Println("Hello. I'm here")

	conf, err := initConfig()
	if err != nil {
		log.Println("Error while parsing config flags: ", err.Error())
		os.Exit(1)
	}

	a := app.NewApp()

	<-a.Start(conf.port, conf.query)
}

type config struct {
	port  string
	query string
}

func initConfig() (config, error) {
	conf := config{}

	flags := flag.NewFlagSet("challenge", flag.ContinueOnError)
	flags.StringVar(&conf.query, "query", defaultQuery, "defines github search query")
	flags.StringVar(&conf.port, "port", defaultPort, "defines server port")

	args := os.Args[1:]
	if err := flags.Parse(args); err != nil {
		return config{}, err
	}
	return conf, nil
}
