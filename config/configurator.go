package config

import (
	"os"
	"time"
)

var (
	TIMEQUERYLAYOUT = "2016-01-02"
	DEFAULTEXPIRETIME = time.Hour*168 //default time is a whole week
)

func init() {
	TIMEQUERYLAYOUT = getEnv("TIME_QUERY_LAYOUT", "2006-01-02")
	//strExpireTime := getEnv("DEFAULT_EXPIRE_TIME", "2")
	//integ, err := strconv.ParseInt(strExpireTime, 10, 64)
	//if err == nil{
	//	DEFAULTEXPIRETIME = time.Duration(time.Duration.Hours(int(integ)))
	//}

}

func getEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return fallback
}
