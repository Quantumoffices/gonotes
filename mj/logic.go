package mj

import (
	"encoding/hex"
	"fmt"
	"math/rand"
	"reflect"
	"sync"
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
func CheckHu(cards [10]byte, specials []byte) (isHu bool, groups []Group) {
	count := len(cards)
	godCount := len(specials)
	if (count+godCount)%3 != 2 {
		return
	}
	//byte to index 转换成牌值存储
	mList := map[byte][10]byte{} //color-value list
	for i := 0; i < count; i++ {
		color := GetCardColor(cards[i])
		value := GetCardValue(cards[i])
		list, exist := mList[color]
		if !exist {
			mList[color] = [MaxCard + 1]byte{value}
			continue
		}
		list[value]++
		mList[color] = list
	}
	doubleCards := []Card{}
	for color, list := range mList {
		for i := 1; i < 10; i++ {
			//1.优先组合可能的将牌
			if list[i] > 1 {
				doubleCard := Card{
					Color: color,
					Value: byte(i),
				}
				doubleCards = append(doubleCards, doubleCard)
			}
			//2.检查是否可333
			if list[i] > 1 {
				list[i] -= 2
				ok, mGroups, lastGodCount := IsCan333(list, godCount, color)
				if !ok {
					return
				}
				godCount = lastGodCount
				groups = append(groups, mGroups...)
			}
		}
	}

	return
}

func Covert2Card(color, value byte) byte {
	str := fmt.Sprintf("%d%d", color, value)
	b, _ := hex.DecodeString(str)
	return b[0]
}

//是否可以组成333形式
//@param list [10]byte 牌值索引
//@param godCount int 百搭牌数量
func IsCan333(list [10]byte, godCount int, color byte) (ok bool, groups []Group, lastGodCount int) {
	for i := 1; i < 10; i++ {
		for list[i] > 0 {
			value := byte(i)
			switch {
			case list[i] >= 3: //优先组碰
				list[i] -= 3
				fmt.Println(list, "去掉3个:", value)
				//card := Covert2Card(color, value)
				card := Card{Color: color, Value: value}
				newGroup := Group{
					Type:  GroupTypePenZi, //碰子
					Cards: []Card{card, card, card},
				}
				groups = append(groups, newGroup)
			case list[i] == 2:
				//212
				//2111
				if godCount > 0 && list[i+1] < 1 ||
					godCount > 0 && i <= 7 && list[i+1] < 2 && list[i+2] < 2 ||
					godCount > 0 && value == MaxCard {
					godCount--
					list[i]++
					continue
				}
				fallthrough
			default:
				//字牌无法组成顺子
				if color >= CardColorZiPai && godCount < 2 {
					return
				}
				//组成顺子
				if i < 8 && list[i+1] > 0 && list[i+2] > 0 {
					list[i]--
					list[i+1]--
					list[i+2]--
					fmt.Println(list, "去掉:", value, "-", value+1, "-", value+2)
					card1 := Card{Color: color, Value: value}
					card2 := Card{Color: color, Value: value + 1}
					card3 := Card{Color: color, Value: value + 2}
					//card1 := Covert2Card(color, value)
					//card2 := Covert2Card(color, value+1)
					//card3 := Covert2Card(color, value+2)
					newGroup := Group{
						Type:  GroupTypeShunZi, //碰子
						Cards: []Card{card1, card2, card3},
					}
					groups = append(groups, newGroup)
				} else if godCount > 0 {
					//89
					if i == 8 && list[i+1] > 0 {
						list[i-1]++
						i--
						godCount--
						continue
					}
					//1/0/0-->组成碰
					if list[i+1] < 0 && list[i+2] < 0 {
						list[i]++
						godCount--
						continue
					}
					//110
					if list[i+1] > 0 && list[i+2] < 1 {
						list[i+2]++
						godCount--
						continue
					}
					//101
					if list[i+2] > 0 && list[i+1] < 1 {
						list[i+1]++
						godCount--
						continue
					}
				} else {
					if value >= MaxCard || godCount < 1 || list[i] > 0 && list[i+1] < 1 {
						return
					}
				}
			}
		}
	}
	ok = true
	lastGodCount = godCount
	return
}

//是否可以组成333形式
//@param list [10]byte 牌值索引
//@param godCount int 百搭牌数量
func Analysis333(list [10]byte, color byte, godCount int) (ok bool, groups []Group) {
	//两组情况
	//优先碰子
	for i := 1; i < 10; i++ {
		for list[i] > 0 {
			value := byte(i)
			switch {
			case list[i] >= 3: //优先组碰
				list[i] -= 3
				fmt.Println(list, "去掉3个:", value)
				card := Card{Color: color, Value: value}
				newGroup := Group{
					Ok:    true,
					Type:  GroupTypePenZi, //碰子
					Cards: []Card{card, card, card},
				}
				groups = append(groups, newGroup)
			case list[i] == 2:
				if color >= CardColorZiPai || i == 9 || list[i+1] < 2 {
					list[i] -= 2
					fmt.Println(list, "去掉2个:", value)
					card := Card{Color: color, Value: value}
					newGroup := Group{
						Ok:    true,
						Type:  GroupTypeDouble, //对子
						Cards: []Card{card, card},
					}
					groups = append(groups, newGroup)
					continue
				}
				fallthrough
			default:
				//字牌无法组成顺子
				if color >= CardColorZiPai {
					return
				}
				//组成顺子
				if i < 8 && list[i+1] > 0 && list[i+2] > 0 {
					list[i]--
					list[i+1]--
					list[i+2]--
					fmt.Println(list, "去掉:", value, "-", value+1, "-", value+2)
					card1 := Card{Color: color, Value: value}
					card2 := Card{Color: color, Value: value + 1}
					card3 := Card{Color: color, Value: value + 2}
					newGroup := Group{
						Ok:    true,
						Type:  GroupTypeShunZi, //碰子
						Cards: []Card{card1, card2, card3},
					}
					groups = append(groups, newGroup)
				} else {
					//11
					if i < 9 && list[i+1] > 0 {
						list[i]--
						list[i+1]--
						//list[i+2]--
						fmt.Println(list, "去掉:", value, "-", value+1, "-", value+2)
						card1 := Card{Color: color, Value: value}
						card2 := Card{Color: color, Value: value + 1}
						//card3 := Card{Color: color, Value: value + 2}
						newGroup := Group{
							Ok:    false,
							Type:  GroupTypeShunZi, //张
							Cards: []Card{card1, card2, Card{InvalidColor, InvalidValue}},
						}
						groups = append(groups, newGroup)
						continue
					}
					//101
					if i < 8 && list[i+2] > 0 {
						list[i]--
						//list[i+1]--
						list[i+2]--
						fmt.Println(list, "去掉:", value, "-", value+1, "-", value+2)
						card1 := Card{Color: color, Value: value}
						//card2 := Card{Color: color, Value: value + 1}
						card3 := Card{Color: color, Value: value + 2}
						newGroup := Group{
							Ok:    false,
							Type:  GroupTypeShunZi, //张
							Cards: []Card{card1, Card{InvalidColor, InvalidValue}, card3},
						}
						groups = append(groups, newGroup)
						continue
					}
					//1/0/0-->单张
					list[i]--
					//list[i+1]--
					//list[i+2]--
					fmt.Println(list, "去掉:", value, "-", value+1, "-", value+2)
					card1 := Card{Color: color, Value: value}
					//card2 := Card{Color: color, Value: value + 1}
					//card3 := Card{Color: color, Value: value + 2}
					newGroup := Group{
						Ok:    true,
						Type:  GroupTypeSingle, //张
						Cards: []Card{card1},
					}
					groups = append(groups, newGroup)
				}
			}
		}
	}
	return
}

func check333(startCard byte, list [MaxCard]byte, godCount int) (ok bool, groups []Group) {
	for curCard := startCard; curCard < MaxCard; curCard++ {
		if list[curCard] >= 3 {
			list[curCard] -= 3
			fmt.Println(list, "去掉3张:", curCard)
			newGroup := Group{
				Ok:    true,
				Type:  GroupTypePenZi, //碰子
				Cards: []byte{curCard, curCard, curCard},
			}
			groups = append(groups, newGroup)
			if ok, otherGroups := check333(curCard, list, godCount); ok {
				groups = append(groups, otherGroups...)
			}
		}
		return
	}
	ok = true
	return
}

func check123(startCard byte, list [MaxCard]byte, godCount int) (ok bool, groups []Group) {
	for curCard := startCard; curCard < MaxCard; curCard++ {
		if list[curCard] > 0 {
			list[curCard] -= 3
			fmt.Println(list, "去掉3张:", curCard)
			newGroup := Group{
				Ok:    true,
				Type:  GroupTypePenZi, //碰子
				Cards: []byte{curCard, curCard, curCard},
			}
			groups = append(groups, newGroup)
			if ok, otherGroups := check333(curCard, list, godCount); ok {
				groups = append(groups, otherGroups...)
			}
		}
		return
	}
	ok = true
	return
}

//优先组碰
func analysis333(startCard byte, list [MaxCard]byte, godCount int) (ok bool, groups []Group) {
	for curCard := startCard; curCard < MaxCard; curCard++ {
		if list[curCard] > 0 {
			switch {
			case list[curCard] >= 3: //优先组碰
				list[curCard] -= 3
				fmt.Println(list, "去掉3张:", curCard)
				newGroup := Group{
					Ok:    true,
					Type:  GroupTypePenZi, //碰子
					Cards: []byte{curCard, curCard, curCard},
				}
				groups = append(groups, newGroup)
				ok, otherGroups := analysis333(curCard, list, godCount)
				if ok {
					groups = append(groups, otherGroups...)
					return
				}
			case list[curCard] == 2:
				//组成对子
				if curCard > 30 || curCard%10 == 9 || list[curCard+1] < 2 {
					list[curCard] -= 2
					fmt.Println(list, "去掉2个:", curCard)
					newGroup := Group{
						Ok:    true,
						Type:  GroupTypeDouble, //对子
						Cards: []byte{curCard, curCard},
					}
					groups = append(groups, newGroup)
					continue
				}
			case list[curCard] == 1:
			}
		}
	}
	return
}

//优先组顺子
//@param list [10]byte 牌值索引
//@param godCount int 百搭牌数量
func Analysis123(list [10]byte, color byte, godCount int) (ok bool, groups []Group) {
	//两组情况
	//优先碰子
	for i := 1; i < 10; i++ {
		for list[i] > 0 {
			value := byte(i)
			switch {
			case list[i] == 1:
				//字牌无法组成顺子
				if color >= CardColorZiPai {
					return
				}
				//组成顺子
				if i < 8 && list[i+1] > 0 && list[i+2] > 0 {
					list[i]--
					list[i+1]--
					list[i+2]--
					fmt.Println(list, "去掉:", value, "-", value+1, "-", value+2)
					card1 := Card{Color: color, Value: value}
					card2 := Card{Color: color, Value: value + 1}
					card3 := Card{Color: color, Value: value + 2}
					newGroup := Group{
						Ok:    true,
						Type:  GroupTypeShunZi, //碰子
						Cards: []Card{card1, card2, card3},
					}
					groups = append(groups, newGroup)
				} else {
					//11
					if i < 9 && list[i+1] > 0 {
						list[i]--
						list[i+1]--
						//list[i+2]--
						fmt.Println(list, "去掉:", value, "-", value+1, "-", value+2)
						card1 := Card{Color: color, Value: value}
						card2 := Card{Color: color, Value: value + 1}
						//card3 := Card{Color: color, Value: value + 2}
						newGroup := Group{
							Ok:    false,
							Type:  GroupTypeShunZi, //张
							Cards: []Card{card1, card2, Card{InvalidColor, InvalidValue}},
						}
						groups = append(groups, newGroup)
						continue
					}
					//101
					if i < 8 && list[i+2] > 0 {
						list[i]--
						//list[i+1]--
						list[i+2]--
						fmt.Println(list, "去掉:", value, "-", value+1, "-", value+2)
						card1 := Card{Color: color, Value: value}
						//card2 := Card{Color: color, Value: value + 1}
						card3 := Card{Color: color, Value: value + 2}
						newGroup := Group{
							Ok:    false,
							Type:  GroupTypeShunZi, //张
							Cards: []Card{card1, Card{InvalidColor, InvalidValue}, card3},
						}
						groups = append(groups, newGroup)
						continue
					}
					//1/0/0-->单张
					list[i]--
					//list[i+1]--
					//list[i+2]--
					fmt.Println(list, "去掉:", value, "-", value+1, "-", value+2)
					card1 := Card{Color: color, Value: value}
					//card2 := Card{Color: color, Value: value + 1}
					//card3 := Card{Color: color, Value: value + 2}
					newGroup := Group{
						Ok:    true,
						Type:  GroupTypeSingle, //张
						Cards: []Card{card1},
					}
					groups = append(groups, newGroup)
				}

			case list[i] == 2:
				if color >= CardColorZiPai || i == 9 || list[i+1] < 2 {
					list[i] -= 2
					fmt.Println(list, "去掉2个:", value)
					card := Card{Color: color, Value: value}
					newGroup := Group{
						Ok:    true,
						Type:  GroupTypeDouble, //对子
						Cards: []Card{card, card},
					}
					groups = append(groups, newGroup)
					continue
				}
			case list[i] >= 3: //优先组碰
				list[i] -= 3
				fmt.Println(list, "去掉3个:", value)
				card := Card{Color: color, Value: value}
				newGroup := Group{
					Ok:    true,
					Type:  GroupTypePenZi, //碰子
					Cards: []Card{card, card, card},
				}
				groups = append(groups, newGroup)

			}
		}
	}
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
