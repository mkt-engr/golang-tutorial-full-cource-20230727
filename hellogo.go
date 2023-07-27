package main

import (
	"fmt"
	"reflect"
	"strconv"
)

var pl = fmt.Println

func main() {
	//キャスト
	//小数型から整数型
	cV1 := 1.5
	cV2 := int(cV1)
	pl(cV2) //1

	// 文字列から数値型
	cV3 := "50000"
	cV4, err := strconv.Atoi(cV3)
	pl(cV4, err, reflect.TypeOf(cV4))

	// 数値から文字列
	cV5 := 5000
	cV6 := strconv.Itoa(cV5)
	pl(cV6)

	// 数値から文字列
	cV7 := "3.14"

	pl(cV6)
	if cV8, err := strconv.ParseFloat(cV7, 64); err == nil {
		pl(cV8)
	}
	cV9 := fmt.Sprintf("%f", 3.14)
	pl(cV9)
}
