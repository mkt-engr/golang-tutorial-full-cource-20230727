package main

import (
	"fmt"
	"strings"
)

var pl = fmt.Println

func main() {
	sV1 := "A word"
	//文字列を置換するメソッドの作成的な
	replacer := strings.NewReplacer("A", "Another")
	sV2 := replacer.Replace(sV1)
	pl(sV2)

	//文字列の長さ出力
	pl("Length : ", len(sV2))
	pl("Contains Another", strings.Contains(sV2, "Another"))

	//指定した文字が何文字目にあるか
	pl("o index :", strings.Index(sV2, "o"))
	// 置換
	pl("Replace :", strings.Replace(sV2, "o", "0", -1))

	//トリミング
	sV3 := "\nSome Words\n"
	sV3 = strings.TrimSpace(sV3)

	//分離
	pl("Split :", strings.Split("a-b-c-d", "-"))

	//大文字小文字変換
	pl("Lower :", strings.ToLower(sV2))
	pl("Upper :", strings.ToUpper(sV2))

	//最初と最後に指定した文字があるか
	pl("Prefix :", strings.HasPrefix("tacocat", "taco"))
	pl("Suffix :", strings.HasSuffix("tacocat", "cat"))

}
