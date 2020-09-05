package main

import (
	"fmt"
	"reflect"
	"testing"
)

func TestIface1(t *testing.T) {
	var (
		data  *int
		eface interface{}
	)
	eface = data
	fmt.Println(data == nil)
	fmt.Println(eface == nil)
	fmt.Println(reflect.TypeOf(data))
	fmt.Println(reflect.ValueOf(data))
}

func TestIface2(t *testing.T) {
	//data := []int{1, 2, 3}
	//printAny((data))
}

func printAny(any []interface{}) {
	fmt.Println(any)
}
