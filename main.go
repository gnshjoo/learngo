package main

import "fmt"

func main() {
	a := 2
	b := &a
	a = 2020
	fmt.Println(*b)
}
