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
	//3332
	cards := []byte{1, 1, 1, 2, 2, 2, 3, 3, 3, 4, 4, 11, 12, 13}
	//411-311
	//cards := []byte{0x01, 0x01, 0x01, 0x01, 0x02, 0x03, 0x07, 0x08, 0x09, 0x07, 0x07}
	//233-1
	//cards := []byte{0x01, 0x01, 0x02, 0x02, 0x02, 0x03, 0x03, 0x03 /*, 0x09*/}
	//godCount := 1
	//2441-1
	//cards := []byte{0x01, 0x01, 0x02, 0x02, 0x02, 0x02, 0x03, 0x03, 0x03, 0x03, 0x04 /*, 0x09*/}
	//godCount := 2
	//洗牌
	//shuffle := Shuffle(BaseCards[:36])
	//cards = shuffle[:11]
	//valueList := [MaxCard + 1]byte{}
	//for _, card := range cards {
	//	v := GetCardValue(card)
	//	valueList[v]++
	//}
	//fmt.Println(valueList)
	//valueList = [10]byte{0, 1, 2, 1, 2, 1, 1, 0, 2, 1}
	//valueList = [10]byte{0, 1, 0, 1, 2, 3, 0, 2, 1, 1}
	hu, groups := CheckHu(cards, []byte{})
	if hu {
		for _, groupList := range groups {
			fmt.Println(groupList)
		}
	}

	//fmt.Println(hu, groups)
}
