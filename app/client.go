package app

import (
	"io/ioutil"
	"log"
	"net/http"
)

type Client struct {
}

func (c Client) RunQuery(query string) {
	queryUrl := "https://api.github.com/search/repositories?q=" + query + "+language:go"
	resp, err := http.Get(queryUrl)
	if err != nil {
		log.Println("Request error: ", err)
		return
	}
	defer resp.Body.Close()

	// write to console
	//io.Copy(os.Stdout, resp.Body)

	ioutil.ReadAll(resp.Body)
}
