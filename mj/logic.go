package mj

import (
	"encoding/hex"
	"fmt"
	"math/rand"
	"time"
)

//麻将一般是136张，有的bai地区增du加了花牌，有144张。
//“万”、“zhi筒”、“条”从一到九，各四dao张，合计108张。
//“东”、“南”、“西”、“北”、“中”、“发”、“白（白板）”各四张，合计28张。一副牌共计136张。加上“春夏秋冬、梅兰菊竹”8张花牌，共144张。
//打麻将时每人拿13张，庄家拿14张直接出牌，其余人摸一张再出牌。

//获取花色
func GetCardColor(card byte) byte {
	//return card & MaskCodeColor >> 4
	return card / 10
}

//获取牌值
func GetCardValue(card byte) byte {
	return card % 10
}

//是否胡牌
//@param cards []byte 普通牌
//@param special []byte 百搭牌
func CheckHu(cards []byte, specials []byte) (isHu bool, groups map[int][]Group) {
	count := len(cards)
	godCount := len(specials)
	if (count+godCount)%3 != 2 {
		return
	}
	//牌数量计算
	list := [MaxCard]byte{}
	for i := 0; i < count; i++ {
		list[cards[i]]++
	}
	//万
	startCard := byte(1)
	fmt.Println(list)
	isHu, groups = multiCheck(Group{}, startCard, list, godCount)

	if isHu {
		for kind, groupList := range groups {
			count = 0
			for i := 0; i < len(groupList); i++ {
				if groupList[i].Type == GroupTypeDouble {
					count++
				}
			}
			//必须是一对
			if count != 1 {
				//去掉不符合的
				delete(groups, kind)
			}
		}
		if len(groups) < 1 {
			isHu = false
		}
	}
	return
}

func Covert2Card(color, value byte) byte {
	str := fmt.Sprintf("%d%d", color, value)
	b, _ := hex.DecodeString(str)
	return b[0]
}

//多路判断
func multiCheck(newGroup Group, curCard byte, list [MaxCard]byte, godCount int) (ok bool, groups map[int][]Group) {
	groups = map[int][]Group{}
	curGroups := []Group{newGroup}
	curKind := 0
	success, otherGroups := check333(curCard, list, godCount)
	if success {
		ok = true
		for _, otherGroupList := range otherGroups {
			curKind++
			if _, exit := groups[curKind]; !exit {
				groups[curKind] = curGroups
			}
			groups[curKind] = append(groups[curKind], otherGroupList...)
		}
	}
	success, otherGroups = check123(curCard, list, godCount)
	if success {
		ok = true
		for _, otherGroupList := range otherGroups {
			curKind++
			if _, exit := groups[curKind]; !exit {
				groups[curKind] = curGroups
			}
			groups[curKind] = append(groups[curKind], otherGroupList...)
		}
	}
	success, otherGroups = check222(curCard, list, godCount)
	if success {
		ok = true
		for _, otherGroupList := range otherGroups {
			curKind++
			if _, exit := groups[curKind]; !exit {
				groups[curKind] = curGroups
			}
			groups[curKind] = append(groups[curKind], otherGroupList...)
		}
	}
	if len(groups) < 1 {
		groups[curKind] = curGroups
	}
	return
}

//组碰子
func check333(startCard byte, list [MaxCard]byte, godCount int) (ok bool, groups map[int][]Group) {
	for curCard := startCard; curCard < MaxCard; curCard++ {
		if list[curCard] > 0 {
			switch {
			case list[curCard] >= 3:
				list[curCard] -= 3
			case list[curCard] == 2 && godCount > 0:
				list[curCard] -= 2
				godCount--
			case list[curCard] == 1 && godCount > 1:
				list[curCard] -= 1
				godCount -= 2
			default:
				return
			}
			newGroup := Group{
				Type:  GroupTypePenZi,
				Cards: []byte{curCard, curCard, curCard},
			}
			ok, groups = multiCheck(newGroup, curCard, list, godCount)
			return
		}
	}
	ok = true
	return
}

//组顺子
func check123(startCard byte, list [MaxCard]byte, godCount int) (ok bool, groups map[int][]Group) {
	for curCard := startCard; curCard < MaxCard; curCard++ {
		if list[curCard] > 0 {
			color := curCard / 10
			//字牌不能组成顺子
			if color > 2 {
				return
			}
			card1 := curCard
			card2 := curCard + 1
			card3 := curCard + 2

			value := curCard % 10
			switch {
			case value < 8 && list[curCard+1] > 0 && list[curCard+2] > 0: //111
				list[card1]--
				list[card2]--
				list[card3]--
			case value < 8 && list[curCard+1] < 1 && list[curCard+2] < 1: //100
				if godCount < 2 {
					return
				}
				godCount -= 2
				list[card1]--
			case value < 8 && list[curCard+1] < 1 && list[curCard+2] > 0: //101
				if godCount < 1 {
					return
				}
				list[card1]--
				//list[card2]--
				list[card3]--
				godCount -= 1
			case value < 8 && list[curCard+1] > 0 && list[curCard+2] < 1: //110
				if godCount < 1 {
					return
				}
				list[card1]--
				list[card2]--
				//list[card3]--
				godCount -= 1
			case value < 9 && list[curCard+1] > 0: //11
				if godCount < 1 {
					return
				}
				card3 = card1 - 1
				list[card1]--
				list[card2]--
				//list[card3]--
				godCount -= 1
			default:
				return
			}
			newGroup := Group{
				Type:  GroupTypeShunZi, //碰子
				Cards: []byte{card1, card2, card3},
			}
			ok, groups = multiCheck(newGroup, curCard, list, godCount)
			return
		}
	}
	ok = true
	return
}

//组对子
func check222(startCard byte, list [MaxCard]byte, godCount int) (ok bool, groups map[int][]Group) {
	for curCard := startCard; curCard < MaxCard; curCard++ {
		if list[curCard] > 0 {
			switch {
			case list[curCard] >= 2:
				list[curCard] -= 2
			case list[curCard] == 1 && godCount > 0:
				list[curCard] -= 1
				godCount--
			default:
				return
			}
			newGroup := Group{
				Type:  GroupTypeDouble, //碰子
				Cards: []byte{curCard, curCard},
			}
			ok, groups = multiCheck(newGroup, curCard, list, godCount)
			return
		}
	}
	ok = true
	return
}

//洗牌
func Shuffle(src []byte) (dst []byte) {
	dst = make([]byte, len(src))
	//随机数生成器，加入时间戳保证每次生成的随机数不一样
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	size := len(src)
	randNums := r.Perm(size)
	for i := 0; i < size; i++ {
		dst[i] = src[randNums[i]]
	}
	return
}
