package main

import (
	"fmt"
	"net/http"

	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func main() {
	http.Handle("/metrics", promhttp.Handler())
	fmt.Println("http://127.0.0.1:6060")
	http.ListenAndServe(":6060", nil)
}
