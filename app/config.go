package app

import (
	"flag"
	"os"
)

const (
	defaultPort  = "8888"
	defaultQuery = "golang"
)

type Config struct {
	Port  string
	Query string
}

func Init() (Config, error) {
	conf := Config{}

	flags := flag.NewFlagSet("challenge", flag.ContinueOnError)
	flags.StringVar(&conf.Query, "query", defaultQuery, "defines github search query")
	flags.StringVar(&conf.Port, "port", defaultPort, "defines web port")

	args := os.Args[1:]
	if err := flags.Parse(args); err != nil {
		return Config{}, err
	}
	return conf, nil
}
