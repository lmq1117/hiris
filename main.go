package main

import "fmt"

func main() {
	add(1, 2)
}

func add(a, b int) {
	fmt.Printf("%d + %d = %d", a, b, a+b)
}
