package main

import (
	"fmt"
)

var pl = fmt.Println

func main() {
	//キャスト
	cV1 := 1.5
	cV2 := int(cV1)
	pl(cV2) //1

	cV3 := "50000"
	cV4, err := strconv(cV3)
	pl(cV4, err)
}
