package generator

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestNewPair(t *testing.T) {
	cur := "USDRUB"
	cur1 := "EURUSD"
	p := NewPair(cur)
	p1 := NewPair(cur1)
	assert.Equal(t, p.name, cur, "Name in Pairs must be equal to USDRUB ")
	assert.NotEqual(t, p, p1, "New pairs USDRUB and EURUSD musnt be same")
}

func TestPairs_Name(t *testing.T) {
	p := Pairs{name: "USDRUB"}
	assert.Equal(t, p.Name(), "USDRUB")
}

func TestPairs_Queue(t *testing.T) {
	p := Pairs{queue: []int{100}}
	assert.Equal(t, p.Queue(), []int{100})
}

func TestPairs_Date(t *testing.T) {
	date := time.Now()
	p := Pairs{date: []time.Time{date}}
	assert.Equal(t, p.Date(), []time.Time{date})
}

//func TestPairs_GenVal(t *testing.T) {
//	p := Pairs{
//		name:  "EURRUB",
//		queue: []int{},
//		date:  []time.Time{},
//	}
//
//	emptyConf := config.Config{
//		CurrPairs: []string{},
//		MaxVal:    0,
//		MinVal:    0,
//	}
//
//	conf := config.Config{
//		CurrPairs: []string{"EURRUB"},
//		MaxVal:    100,
//		MinVal:    1,
//	}
//	confi := config.New()
//
//	assert.Equal(t, confi, &emptyConf)
//
//}
