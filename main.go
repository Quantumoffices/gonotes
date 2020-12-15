package main

import (
	"fmt"
	"os"
	"reflect"
	"regexp"
	"unsafe"
)

type A struct {
	S *string
}

func (f *A) String() string {
	return *f.S
}

type ATrick struct {
	S unsafe.Pointer
}

func (f *ATrick) String() string {
	return *(*string)(f.S)
}

func NewA(s string) A {
	return A{S: &s}
}

func NewATrick(s string) ATrick {
	return ATrick{S: noescape(unsafe.Pointer(&s))}
}

func noescape(p unsafe.Pointer) unsafe.Pointer {
	x := uintptr(p)
	return unsafe.Pointer(x ^ 0)
}

type Equaler interface {
	Equal(Equaler) bool
}
type T int

func (t T) Equal(u T) bool { return t == u } // does not satisfy Equaler

func main() {
	var sl []int
	fmt.Println(sl == nil)
	fmt.Println(fmt.Sprintf("%#v", sl))
	fmt.Println(fmt.Sprintf("%#t", sl))
	fmt.Println(reflect.ValueOf(sl))
	fmt.Println(reflect.TypeOf(sl))
	sl = append(sl, 1)
	fmt.Println(sl)
	os.Open()
}

func Bad() (err error) {
	return err
}

//验证邮箱
func VerifyEmailAddr(addr string) bool {
	exp := `\w+([-+.]\w+)*@\w+([-.]\w+)*\.\w+([-.]\w+)*`
	re := regexp.MustCompile(exp)
	return re.MatchString(addr)
}
