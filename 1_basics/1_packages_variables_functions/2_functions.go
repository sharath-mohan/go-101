package main

import (
	"fmt"
)

func add(x int ,y int) int {
	return x+y
}
func sub(x,y int) int {
	return x-y
}

func main(){
	fmt.Println(add(100,200))
	fmt.Println(sub(400,200))
}