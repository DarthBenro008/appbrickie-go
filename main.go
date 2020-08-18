package main

import (
	"appbrickie/api"
	"appbrickie/bot"
	"sync"
)

var wg sync.WaitGroup

func main() {
	wg.Add(2)
	go bot.InitialiseBot()
	go api.InitialiseApi()
	wg.Wait()
}
