package main

import "fmt"

func main() {
	src := make([]int, 0)
	add(src)
	fmt.Println(src)
	fmt.Println(fmt.Sprintf("%p", src))
}

func add(src []int) {

	fmt.Println(fmt.Sprintf("%p", src))

	for i := 0; i < 10; i++ {
		src = append(src, i)
		fmt.Println(fmt.Sprintf("%p", src))
	}
}
