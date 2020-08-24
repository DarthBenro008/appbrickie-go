package main

import (
	"appbrickie/db"
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
	wg.Add(1)
	//go bot.InitialiseBot()
	//go api.InitialiseApi()
	go db.InitialiseDatabase()
	wg.Wait()
}
