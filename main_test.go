package main

import (
	"fmt"
	"testing"
	"time"
)

func TestIface1(t *testing.T) {
	fmt.Println("20006/01/02", time.Now().String())
}

func TestIface2(t *testing.T) {
	//data := []int{1, 2, 3}
	//printAny((data))
}

func printAny(any []interface{}) {
	fmt.Println(any)
}
