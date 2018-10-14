package config

import "os"

var (
	TIMEQUERYLAYOUT = ""
)

func init() {
	TIMEQUERYLAYOUT = getEnv("TIMEQUERYLAYOUT", "2006-01-02")
}

func getEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return fallback
}