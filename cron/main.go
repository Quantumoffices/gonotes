package main

import (
	"fmt"
	"github.com/robfig/cron/v3"
	"github.com/sirupsen/logrus"
	"time"
)

//Field name   | Mandatory? | Allowed values  | Allowed special characters
//----------   | ---------- | --------------  | --------------------------
//Minutes      | Yes        | 0-59            | * / , -
//Hours        | Yes        | 0-23            | * / , -
//Day of month | Yes        | 1-31            | * / , - ?
//Month        | Yes        | 1-12 or JAN-DEC | * / , -
//Day of week  | Yes        | 0-6 or SUN-SAT  | * / , - ?
func main() {
	// Seconds field, optional
	//cronSele:=cron.New(cron.WithParser(cron.SecondOptional | cron.Minute | cron.Hour | cron.Dom | cron.Month | cron.Dow | cron.Descriptor))
	//fmt.Println(cronSelf)
	logrus.Println("wwww")
	// Seconds field, optional
	// Seconds field, required
	c := cron.New(cron.WithSeconds())
	//c.AddFunc("* * * * * *", func() { fmt.Println("Every second") })
	c.Start()
	c.AddFunc("@every 1s", func() {
		fmt.Println("Every second")
	})
	cron.NewChain()
	<-time.After(time.Second * 5)
}
