package main

import (
	"PriceGen/api"
	"PriceGen/config"
	"PriceGen/service/generator"
	"github.com/joho/godotenv"
	"log"
	"sync"
)

func init() {
	if err := godotenv.Load(); err != nil {
		log.Print("No .env file found")
	}
}

func main() {
	wg := &sync.WaitGroup{}
	conf := config.New()
	for _, curr := range conf.CurrPairs {
		wg.Add(1)
		go func(name string) {
			p := generator.NewPair(name)
			p.GenVal(wg, conf)
		}(curr)
	}
	api.StartServer()
	wg.Wait()
}
