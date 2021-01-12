package main

import (
	"fmt"
	"strings"
)

func main() {
	str := "I am a student."
	splits := strings.Split(str, " ")
	fmt.Println(splits)

}

func IsShunZi(cards []int) bool {
	if len(cards) < 5 {
		return false
	}
	//排序
	for i := 0; i < len(cards)-1; i++ {
		for j := i + 1; j < len(cards); j++ {
			if cards[j] < cards[i] {
				cards[i], cards[j] = cards[j], cards[i]
			}
		}
	}
	fmt.Println(cards)
	//
	for i := 0; i < len(cards)-1; i++ {
		if cards[i] == 0 {
			return false
		}
		if cards[i]+1 != cards[i+1] {
			return false
		}
	}
	return true
}
