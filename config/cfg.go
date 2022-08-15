package config

import (
	"os"
	"strconv"
	"strings"
)

type Config struct {
	CurrPairs []string
	MaxVal    int
	MinVal    int
	MaxDelay  int
	MinDelay  int
}

func New() *Config {
	return &Config{
		CurrPairs: getEnvSlice("CURR_PAIRS", []string{}, ","),
		MaxVal:    getEnvInt("MAX_VALUE", 0),
		MinVal:    getEnvInt("MIN_VALUE", 0),
		MaxDelay:  getEnvInt("MAX_DELAY", 1),
		MinDelay:  getEnvInt("MIN_DELAY", 0),
	}
}

func getEnvSlice(name string, defaultVal []string, sep string) []string {
	valStr := getEnv(name, "")

	if valStr == "" {
		return defaultVal
	}

	val := strings.Split(valStr, sep)
	return val
}

func getEnv(key string, defaultVal string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return defaultVal
}

func getEnvInt(name string, defaultVal int) int {
	valueStr := getEnv(name, "")
	if value, err := strconv.Atoi(valueStr); err == nil {
		return value
	}
	return defaultVal
}
