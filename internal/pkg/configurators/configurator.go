package configurators

import (
	"os"
	"strconv"
	"time"
)

 
type HTTPConfigurator struct {
	Listen string
}
type RestConfigurator struct {
	TimeQueryLayout string
}
type PoppinsConfigurator struct {
	Expiretime time.Duration
}

func NewHTTPConfigurator() HTTPConfigurator{
	config := HTTPConfigurator{}
	config.Listen = getEnv("REST_PORT","0.0.0.0:8080")
	return config
}
func NewRestConfigurator() RestConfigurator{
	config := RestConfigurator{}
	config.TimeQueryLayout = getEnv("REST_QUERYPARAM_TIME_LAYOUT","2006-02-01")
	return config
}
func NewPoppinsConfigurator() PoppinsConfigurator{
	config := PoppinsConfigurator{}
	hours, err := time.ParseDuration(getEnv("POPPINS_EXPIRE_TIME","168h"))
	if err != nil{
		hours= time.Hour*168
	}
	config.Expiretime = hours
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
