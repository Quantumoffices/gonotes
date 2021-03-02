package main

import (
	"fmt"
	"math/rand"
	"time"
)

//添加好友的方式
const (
	AddContactByGroup1 = iota + 1 //群内添加
	AddContactByGroup             //群内添加

	Weigh_Banker_Win     = 0x01
	Weight_Player_Win    = 0x02
	Weight_Draw          = 0x04
	Weight_Banker_Double = 0x08
	Weight_Player_Double = 0x10
)

func main() {
	for {
		time.Sleep(time.Second)
		fmt.Println(time.Now().Second())
	}
}

const (
	RAND_NUM   = 0 // 纯数字
	RAND_LOWER = 1 // 小写字母
	RAND_UPPER = 2 // 大写字母
	RAND_ALL   = 3 // 数字、大小写字母
)

var count int64

//生成随机码
func GenRandCode(inviteCodeLen int, targetKind int) string {
	count++
	ikind, kinds, result := targetKind, [][]int{[]int{10, 48}, []int{26, 97}, []int{26, 65}}, make([]byte, inviteCodeLen)
	isAll := targetKind > 2 || targetKind < 0
	rand.Seed(count)
	for i := 0; i < inviteCodeLen; i++ {
		if isAll {
			ikind = rand.Intn(3)
		}
		scope, base := kinds[ikind][0], kinds[ikind][1]
		result[i] = uint8(base + rand.Intn(scope))
	}
	return string(result)
}
