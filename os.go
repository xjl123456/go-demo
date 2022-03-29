package main

import (
	"fmt"
	"os"
)

func main2() {
	aa := "D:\\desktop\\1231231\\functions\\Knative\\hello-world-go"
	bb := "D:\\desktop\\1231232"
	err1 := os.RemoveAll(aa)
	if err1 != nil {
		fmt.Println("-----")
		fmt.Println(err1)
	}
	err := os.Rename(aa, bb)
	if err != nil {
		fmt.Println(err)
	}
}
