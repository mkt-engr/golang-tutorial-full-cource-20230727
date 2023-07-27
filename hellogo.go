package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

var pl = fmt.Println

func main() {
	pl("What is your name?")
	reader := bufio.NewReader(os.Stdin)
	name, err := reader.ReadString('\n')
	if err == nil {
		pl("Hello", name)
	} else {
		log.Fatal(err)
	}

	var vName string = "Makito"
	var v_name string = "Makito"
	var v1, v2 = 1.2, 3.4
	var v3 = "hello"
	v4 := 2.4 //:=はvarの省略記法　つまりvar v4 = 2.4と同じ。
	// v5 = 1 varか:=はつける必要ある

}
