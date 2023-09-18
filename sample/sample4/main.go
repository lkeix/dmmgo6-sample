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
	ch := make(chan bool, 1)
	go func() {
		defer fuga(msg)
		msg = "fugafuga"
		ch <- true
	}()
	<-ch
}
