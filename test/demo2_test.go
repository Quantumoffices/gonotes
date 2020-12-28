package test

import (
	"fmt"
	"testing"
)

func TestHello(t *testing.T) {
	fmt.Println("hello world")
}

func TestAdd(t *testing.T) {
	a := 10
	b := 11
	fmt.Println("a+b=", a+b)
}

// 斐波那契数列-多种实现方式
// 求出第n个数的值
//递归方式
//func Fibonacci(n int) int64 {
//	if n < 2 {
//		return n
//	}
//	return Fibonacci(n-1) + Fibonacci(n-2)
//}

////闭包
//func Fibonacci(n int) (nums []int) {
//	a, b := 0, 1
//	f := func() int {
//		a, b = b, a+b
//		return a
//	}
//	nums = append(nums, 0)
//	for i := 0; i < n; i++ {
//		//fmt.Println(f())
//		nums = append(nums, f())
//	}
//	return
//}

//轮询
func Fibonacci(n int) (nums []int) {
	a, b := 0, 1
	nums = append(nums, 0)
	for i := 0; i < n; i++ {
		a, b = b, a+b
		nums = append(nums, a)
	}
	return
}

func TestFibonacci(t *testing.T) {
	fmt.Println(Fibonacci(20))
}

// 测试参数为20的性能
func BenchmarkFibonacci20(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Fibonacci(20)
	}
}
