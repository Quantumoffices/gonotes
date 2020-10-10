package mj

const (
	MaxCountRepository      = 136 //库存数量
	MaxCountHandCard        = 17  //手牌数量
	MaxChair                = 4   //最大椅子数量
	MaxCard            byte = 40  //最大牌

	//MaskCodeColor byte = f0 //花色掩码
	//MaskCodeValue      = 0f //牌值掩码
)
const (
	InvalidColor byte = 0xff //无效花色
	InvalidValue      = 0xff //无效牌值
	InvalidCard       = 0xff //无效牌
)

//牌类
const (
	CardColorWan    byte = iota //万
	CardColorTong               //筒
	CardColorTiao               //条
	CardColorZiPai              //字牌
	CardColorHuaPai             //花牌
)

//组合类型
const (
	GroupTypeSingle byte = iota //单张
	GroupTypeDouble             //对子
	GroupTypePenZi              //碰子
	GroupTypeShunZi             //顺子
)

type Card struct {
	Color byte
	Value byte
}

type Group struct {
	Type  byte
	Cards []byte
}

//牌数据
var BaseCards = []byte{
	//36
	1, 2, 3, 4, 5, 6, 7, 8, 9, //1~9万
	1, 2, 3, 4, 5, 6, 7, 8, 9, //1~9万
	1, 2, 3, 4, 5, 6, 7, 8, 9, //1~9万
	1, 2, 3, 4, 5, 6, 7, 8, 9, //1~9万
	//36
	11, 12, 13, 14, 15, 16, 17, 18, 19, //1~9筒
	11, 12, 13, 14, 15, 16, 17, 18, 19, //1~9筒
	11, 12, 13, 14, 15, 16, 17, 18, 19, //1~9筒
	11, 12, 13, 14, 15, 16, 17, 18, 19, //1~9筒
	//36
	21, 22, 23, 24, 25, 26, 27, 28, 29, //1~9条
	21, 22, 23, 24, 25, 26, 27, 28, 29, //1~9条
	21, 22, 23, 24, 25, 26, 27, 28, 29, //1~9条
	21, 22, 23, 24, 25, 26, 27, 28, 29, //1~9条
	//28
	31, 32, 33, 34, 35, 36, 37, //东南西北中发白（字牌）
	31, 32, 33, 34, 35, 36, 37, //东南西北中发白（字牌）
	31, 32, 33, 34, 35, 36, 37, //东南西北中发白（字牌）
	31, 32, 33, 34, 35, 36, 37, //东南西北中发白（字牌）
}

//春夏秋冬梅兰竹菊（花牌） 8
var SpecialCards = []byte{41, 42, 43, 44, 45, 46, 47, 48}
