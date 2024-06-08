package main

import "fmt"

func recov() {
	str := recover()
	fmt.Println(str)
}

func main() {
	defer recov () 
	panic("PANIC")
}
