package main

import (
	"fmt"
)

var hoge = "hoge"

func fuga(msg string) {
	fmt.Println(msg)
}

func main() {
	msg := "hogehoge"
	defer fuga(msg)
	msg = "fugafuga"
}
