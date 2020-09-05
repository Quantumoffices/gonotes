package main

import (
	"fmt"
	"runtime"
	"time"
)

func main() {
	runtime.GOMAXPROCS(1)
	for i := 0; i < 500000; i++ {
		//time.Sleep(time.Millisecond)
		go func() {
			j := i
			fmt.Println(j)
		}()
	}

	time.Sleep(time.Second)
	//time.Sleep(time.Second)
	//for i := 0; i < 10; i++ {
	//	go func() {
	//		//time.Sleep(time.Millisecond * 10)
	//		//j := i
	//		//fmt.Println(j)
	//		fmt.Println(i)
	//	}()
	//}

	//var arr []int = []int{1, 2, 3}
	//arrUinrptr := uintptr(unsafe.Pointer(&arr[1]))
	//arr2 := (*int)(unsafe.Pointer(arrUinrptr + 1))
	//fmt.Println(fmt.Sprintf("%v", *arr2))
	//
	//var a = "nnnn"
	//fmt.Println(a)
	//var b = (*int)(unsafe.Pointer(uintptr(unsafe.Pointer(&a)) + 8))
	//// 按照 stringStruct 结构，把 a 地址偏移 int 的长度位，得到 len 字段地址
	//// 这里我的电脑是 64 位，而系统寻址以一个在节为单位，所以 +8
	//fmt.Println(*b) // 这里输出的是 a 的长度 4
}
