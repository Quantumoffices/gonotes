package main

import (
	"fmt"
	"math"
	"math/big"
)

func main() {
	value := float64(110)
	decimal := int64(6)
	//这是处理位数的代码段
	amount := big.NewFloat(float64(value))
	fmt.Println("amount=", amount)
	tenDecimal := big.NewFloat(math.Pow(10, float64(decimal)))
	fmt.Println("tenDecimal=", tenDecimal)
	convertAmount, _ := new(big.Float).Mul(tenDecimal, amount).Int(&big.Int{})
	fmt.Println("convertAmount=", convertAmount)
	//log.Println("转账数量：", convertAmount)
	has := fmt.Sprintf("%x", convertAmount) //格式化数据
	fmt.Println(has)
	afterconvertAmount, err := ParseBigInt("0x68e7780")

	if err != nil {
		fmt.Println(err)
		return
	}

	convertAmountFloat := new(big.Float).SetInt(&afterconvertAmount)
	fmt.Println("convertAmountFloat=", convertAmountFloat)

	fmt.Println(new(big.Float).Quo(convertAmountFloat, tenDecimal))

}

// 将十六进制字符串值解析为big.Int
func ParseBigInt(value string) (big.Int, error) {
	i := big.Int{}
	_, err := fmt.Sscan(value, &i)
	return i, err
}
