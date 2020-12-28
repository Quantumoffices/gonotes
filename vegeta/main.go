package main

import (
	"fmt"
	vegeta "github.com/tsenart/vegeta/v12/lib"
	"time"
)

func main() {
	// 1. 速率&压测时长
	rate := vegeta.Rate{Freq: 100, Per: time.Second}
	duration := 1 * time.Second
	//2.压测目标
	target := vegeta.NewStaticTargeter(vegeta.Target{
		Method: "GET",
		URL:    "https://www.baidu.com",
	})
	// 3. 启动压测并收集结果
	attacker := vegeta.NewAttacker()
	var metrics vegeta.Metrics
	count := 0
	for res := range attacker.Attack(target, rate, duration, "Big Bang!") {
		count++
		fmt.Println("resp:", count, res.Code)
		metrics.Add(res)
	}
	metrics.Close()

	fmt.Printf("99th percentile: %#v\n", metrics.StatusCodes)
}
