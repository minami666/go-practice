package main

import (
	"fmt"
)

func main() {
	fmt.Println(add("元気", "勇気"))
}

func add(x, y string) string {
	fmt.Printf("%sと%sこそがすべて\n", x, y)
	return x + y
}
