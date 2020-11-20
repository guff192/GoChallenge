package app

import (
	"io/ioutil"
	"log"
	"net/http"
)

type Client struct {
}

func (c Client) RunQuery() {
	resp, err := http.Get("https://api.github.com/search/repositories?q=golang")
	if err != nil {
		log.Println("Request error: ", err)
		return
	}
	defer resp.Body.Close()

	// write to console
	//io.Copy(os.Stdout, resp.Body)

	ioutil.ReadAll(resp.Body)
}
