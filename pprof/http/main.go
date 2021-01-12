package main

import (
	"fmt"
	"net/http"
	_ "net/http/pprof"
	"runtime"
	"sync"
	"time"
)

//func index(ctx *fasthttp.RequestCtx) {
//	time.Sleep(200 * time.Millisecond)
//	fmt.Fprintf(ctx.Response.BodyWriter(), "aaa")
//}

//参考https://blog.csdn.net/raogeeg/article/details/82754001
//http://127.0.0.1:8899/debug/pprof/
//go tool pprof http://127.0.0.1:8899/debug/pprof/heap?debug=1
func main() {
	runtime.GOMAXPROCS(4)
	go func() {
		fmt.Println("http://127.0.0.1:8899/debug/pprof/")
		http.ListenAndServe("0.0.0.0:8899", nil)
	}()
	w := sync.WaitGroup{}
	w.Add(1000)
	for i := 0; i < 100; i++ {
		time.Sleep(time.Second * 5)
		go func() {
			var array [10000000]int64
			var array2 [10000000]int64
			fmt.Println("run goroutine-", i, array[0], array2[0])
			//var ch chan int
			//ch <- 10
			lock := sync.Mutex{}
			lock.Lock()
			lock.Lock()
			w.Done()
		}()
	}
	w.Wait()
}
