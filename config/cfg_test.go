package config

import (
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

func TestGetEnv(t *testing.T) {
	err := os.Setenv("MIN_VALUE", "1")
	if err != nil {
		t.Fatalf("error in set env value")
	}
	l, ok := os.LookupEnv("MIN_VALUE")
	if ok != true {
		t.Fatalf("error in get .env")
	}
	assert.Equal(t, getEnv("MIN_VALUE", ""), l)
	assert.Equal(t, getEnv("MAX_VALUE", ""), "")

}

func TestGetEnvInt(t *testing.T) {
	err := os.Setenv("MIN_VALUE", "1")
	if err != nil {
		t.Fatalf("error in set env value")
	}
	assert.Equal(t, getEnvInt("MIN_VALUE", 0), 1)
	assert.Equal(t, getEnvInt("MIN_VALUE", 0), 1)
}

func TestGetEnvSlice(t *testing.T) {
	err := os.Setenv("CURR_PAIRS", "EURUSD,USDRUB,EURRUB,USDJPY,EURJPY")
	if err != nil {
		t.Fatalf("error in set env value")
	}
	assert.Equal(t, getEnvSlice("CURR_PAIRS", []string{}, ","), []string{"EURUSD", "USDRUB", "EURRUB", "USDJPY", "EURJPY"})
	assert.Equal(t, getEnvSlice("CURRRR_PAIRS", []string{}, ","), []string{})
}

func TestNewConf(t *testing.T) {
	err := os.Setenv("CURR_PAIRS", "")
	if err != nil {
		t.Fatalf("error in set env value")
	}
	err = os.Setenv("MIN_VALUE", "0")
	if err != nil {
		t.Fatalf("error in set env value")
	}
	c := New()
	nc := Config{
		CurrPairs: []string{},
		MaxVal:    0,
		MinVal:    0,
		MaxDelay:  1,
		MinDelay:  0,
	}
	assert.Equal(t, c, &nc)
}
