package main

import "C"
import (
	"fmt"
	"os"
)

func main() {
	test3()
}
func test4() {
	b := initAccount(os.Args[6], os.Args[5], os.Args[4])
	fmt.Printf("unlock:%t\n", b)
	if !b {
		panic("unlock failed")
	}
	b = startService(os.Args[1], os.Args[2], os.Args[3])
	if !b {
		panic("service failed")
	}
	<-make(chan struct{})
}
func test1() {
	b := startService(os.Args[1], os.Args[2], os.Args[3])
	if !b {
		panic("failed")
	}
	<-make(chan struct{})
}
func test2() {
	a, c := createAccount(os.Args[1])
	fmt.Println(a)
	fmt.Println(c)
}
func test3() {
	b := initAccount(os.Args[1], os.Args[2], os.Args[3])
	fmt.Printf("unlock:%t\n", b)
}
