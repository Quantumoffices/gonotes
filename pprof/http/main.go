package main

import (
	"fmt"
	"github.com/valyala/fasthttp"
	"log"
	"net/http"
	_ "net/http/pprof"
	"runtime"
	"time"
)

func index(ctx *fasthttp.RequestCtx) {
	time.Sleep(200 * time.Millisecond)
	fmt.Fprintf(ctx.Response.BodyWriter(), "aaa")
}

//参考https://blog.csdn.net/raogeeg/article/details/82754001
//http://127.0.0.1:8899/debug/pprof/
func main() {
	runtime.GOMAXPROCS(4)

	log.Println("fasthttp")

	m := func(ctx *fasthttp.RequestCtx) {
		switch string(ctx.Path()) {
		case "":
			index(ctx)
		default:
			ctx.Error("not found", fasthttp.StatusNotFound)
		}
	}
	go func() {
		fmt.Println("http://127.0.0.1:8899/debug/pprof/")
		http.ListenAndServe("0.0.0.0:8899", nil)
	}()
	addr := ":8083"
	fasthttp.ListenAndServe(addr, m)

}
