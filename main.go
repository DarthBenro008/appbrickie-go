package main

import (
	"appbrickie/api"
	"appbrickie/bot"
	"github.com/joho/godotenv"
	"log"
	"sync"
)

var wg sync.WaitGroup

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Println("Error Loading env File")
	}
	wg.Add(2)
	go bot.InitialiseBot()
	go api.InitialiseApi()
	wg.Wait()
}
