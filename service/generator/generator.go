package generator

import (
	"PriceGen/api"
	"PriceGen/config"
	"fmt"
	"math/rand"
	"sync"
	"time"
)

type Pairs struct {
	name  string
	queue []int
	date  []time.Time
}

func (p *Pairs) Name() string {
	return p.name
}

func (p *Pairs) Queue() []int {
	return p.queue
}

func (p *Pairs) Date() []time.Time {
	return p.date
}

func NewPair(curr string) *Pairs {
	return &Pairs{name: curr}
}

func (p *Pairs) GenVal(wg *sync.WaitGroup, conf *config.Config) {
	defer wg.Done()
	for {
		if len(p.queue) != 100 {
			p.queue = append(p.queue, rand.Intn(conf.MaxVal)+conf.MinVal)
			p.date = append(p.date, time.Now())
			time.Sleep(time.Duration(rand.Intn(conf.MaxDelay)+conf.MinDelay) * time.Second)
			switch p.name {
			case "EURUSD":
				{
					api.MapCurr["EURUSD"] = p.queue
					api.MapDate["EURUSD"] = p.date
				}
			case "USDRUB":
				{
					api.MapCurr["USDRUB"] = p.queue
					api.MapDate["USDRUB"] = p.date
				}
			case "EURRUB":
				{
					api.MapCurr["EURRUB"] = p.queue
					api.MapDate["EURRUB"] = p.date
				}
			case "USDJPY":
				{
					api.MapCurr["USDJPY"] = p.queue
					api.MapDate["USDJPY"] = p.date
				}
			case "EURJPY":
				{
					api.MapCurr["EURJPY"] = p.queue
					api.MapDate["EURJPY"] = p.date
				}
			default:
				fmt.Println("err")

			}
		} else {
			p.queue = p.queue[1:]
			p.date = p.date[1:]
			continue
		}
	}
}
