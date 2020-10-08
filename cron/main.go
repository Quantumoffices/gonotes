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
//Entry                  | Description                                | Equivalent To
//-----                  | -----------                                | -------------
//@yearly (or @annually) | Run once a year, midnight, Jan. 1st        | 0 0 1 1 *
//@monthly               | Run once a month, midnight, first of month | 0 0 1 * *
//@weekly                | Run once a week, midnight between Sat/Sun  | 0 0 * * 0
//@daily (or @midnight)  | Run once a day, midnight                   | 0 0 * * *
//@hourly                | Run once an hour, beginning of hour        | 0 * * * *
//@every 1h30m10s
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
