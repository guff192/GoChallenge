package web

import (
	"GoChallenge/app/util"
	"time"
)

var (
	GithubResponseChannel = make(chan util.GithubResponse, 1)
	githubResp            util.GithubResponse
	timeElapsed           bool
	delayer               *time.Ticker
	wait                  bool
)
