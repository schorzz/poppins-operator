package configurators

import (
	"os"
	"strconv"
	"time"
)


type Config struct {
	SocketAdress 			string
	HttpTimeQueryLayout 	string
	NamespaceExpTime		time.Duration
}

func New(config Config)Config {
	config.SocketAdress = getEnv("REST_PORT","0.0.0.0:8080")
	config.HttpTimeQueryLayout = getEnv("REST_QUERYPARAM_TIME_LAYOUT","2006-02-01")
	hours, err := time.ParseDuration(getEnv("POPPINS_EXPIRE_TIME","168h"))
	if err != nil{
		hours= time.Hour*168
	}
	config.NamespaceExpTime = hours
	return config
}

func getEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return fallback
}
func getEnvI(key string, val int) int{
	envVal := getEnv(key, "")
	if envVal != ""{
		if s, err := strconv.ParseInt(envVal,10,32); err==nil{
			return int(s)
		}
	}
	return val

}
