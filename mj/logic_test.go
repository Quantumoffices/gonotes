package mj

import (
	"fmt"
	"testing"
)

func TestCovert2Card(t *testing.T) {
	card := Covert2Card(0, 1)
	fmt.Println(card)
}

func TestCheckHu(t *testing.T) {
	//cards := []byte{0x01, 0x02, 0x03, 0x04, 0x05, 0x06, 0x07, 0x08, 0x09}
	//411-0
	//cards := []byte{0x01, 0x01, 0x01, 0x01, 0x02, 0x03, 0x07, 0x08, 0x09}
	//233-1
	//cards := []byte{0x01, 0x01, 0x02, 0x02, 0x02, 0x03, 0x03, 0x03 /*, 0x09*/}
	//godCount := 1
	//2441-1
	cards := []byte{0x01, 0x01, 0x02, 0x02, 0x02, 0x02, 0x03, 0x03, 0x03, 0x03, 0x04 /*, 0x09*/}
	godCount := 2
	//洗牌
	shuffle := Shuffle(BaseCards[:36])
	cards = shuffle[:11]
	valueList := [MaxCard + 1]byte{}
	for _, card := range cards {
		v := GetCardValue(card)
		valueList[v]++
	}
	fmt.Println(valueList)
	//valueList = [10]byte{0, 1, 2, 1, 2, 1, 1, 0, 2, 1}
	//valueList = [10]byte{0, 1, 0, 1, 2, 3, 0, 2, 1, 1}
	can333, groups, count := IsCan333(valueList, godCount, 0)
	fmt.Println(can333, groups, count)
}
