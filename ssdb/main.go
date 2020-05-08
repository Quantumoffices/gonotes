package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"reflect"
)

type parnt struct {
	Name string
}
type children struct {
	parnt
	Age int
}

//登录
type LoginReq struct {
	//Os         string `json:"os"`          //系统
	Platform   string `json:"platform"`    //注册平台 wx/ali/bd/hw/oppo/vivo/ios/act e.g.
	PlatformID string `json:"platform_id"` //平台id
}
type a struct {
}
type Hi struct {
	content string
}

func (a) Say(c Hi) (string, error) {
	fmt.Println(c.content)
	return "reply", errors.New("bb")
}

func main() {

	a := a{}
	method := reflect.ValueOf(a).MethodByName("Say")
	fmt.Println(method.Kind() == reflect.Invalid)
	result1 := method.Call([]reflect.Value{reflect.ValueOf(Hi{"okok"})})
	if len(result1) > 0 {
		fmt.Println(result1[0].Interface())
		fmt.Println(result1[1].Interface().(error).Error())
	}
	return

	data, err := json.Marshal([]byte("fjkaljngkl呵呵"))
	if err != nil {
		panic(err)
	}
	fmt.Println(string(data))

	//n, _ := strconv.ParseFloat("", 4)
	//fmt.Println(n)
	//return
	//xiaomi := LoginReq{
	//	Platform:   "oter",
	//	PlatformID: "123456",
	//}
	//bs, _ := json.Marshal(xiaomi)
	//fmt.Println(string(bs))
	//var a interface{}
	//b := false
	//a = b
	//c := a.(bool)
	//fmt.Println(c)
	//v, _ := strconv.ParseFloat("0.32", 4)
	//fmt.Println(v)
}
