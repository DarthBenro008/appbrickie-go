package main

import (
	"sync"
	"tele-go/api"
	"tele-go/bot"
)

var wg sync.WaitGroup

func main() {
	wg.Add(2)
	go bot.InitialiseBot()
	go api.InitialiseApi()
	wg.Wait()
}
