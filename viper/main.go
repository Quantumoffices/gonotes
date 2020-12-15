package main

import (
	"fmt"
	"github.com/spf13/viper"
	"time"
)

func main() {
	viper.AutomaticEnv()
	fmt.Println(viper.GetString("ENV"))
	time.Sleep(time.Minute)
}
