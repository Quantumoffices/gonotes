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
	//3332123
	cards := []byte{1, 1, 1, 2, 2, 2, 3, 3, 3, 4, 4, 11, 12, 13}
	hu, groups := CheckHu(cards, []byte{})
	if hu {
		for _, groupList := range groups {
			fmt.Println(groupList)
		}
	}
}
