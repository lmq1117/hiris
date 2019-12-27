package main

import (
	"fmt"
	"time"
)

func main() {
	go add(1, 2)
	time.Sleep(1e9)
}

func add(a, b int) {
	fmt.Printf("%d + %d = %d", a, b, a+b)
}
