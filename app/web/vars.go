package web

import (
	"GoChallenge/app/util"
	"time"
)

var (
	GithubResponseChannel = make(chan util.GithubResponse, 1)
	Next                  <-chan time.Time
	GithubResp            util.GithubResponse
	TimeElapsed           = false
	ParseFinished         = false
)
