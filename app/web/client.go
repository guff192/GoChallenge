package web

import (
	"GoChallenge/app/util"
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
		log.Println("Request error: ", err.Error())
		return
	}
	defer resp.Body.Close()

	// write to console
	//io.Copy(os.Stdout, resp.Body)

	bytesResp, _ := ioutil.ReadAll(resp.Body)
	jsonResp, err := util.ParseResponse(bytesResp)
	if err != nil {
		log.Println("Response parsing error: ", err.Error())
		return
	}
	GithubResponseChannel <- *jsonResp
}
