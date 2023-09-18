package main

import (
	"fmt"
)

func fuga(msg string) {
	fmt.Println("fuga")
}

func main() {
	msg := "hogehoge"
	if true {
		defer fuga(msg)
	}
	msg = "fugafuga"
}
