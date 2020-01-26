package main

import (
	"fmt"
)

func main() {
	fmt.Println(add(42, 13))
}

func add(x, y int) int {
	fmt.Printf("%dと%dを使って足し算をしよう\n", x, y)
	return x + y
}
