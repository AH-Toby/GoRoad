package main

import "fmt"

func main() {
	f1 := func(x,y int ) {
		fmt.Println(x+y)
	}
	f1(10,20)
}