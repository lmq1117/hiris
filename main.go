package main

import (
	"fmt"
	"strings"
	"time"
)

func main() {
	go add(1, 2)
	time.Sleep(1e9)
	fmt.Println(strings.Join([]string{"aa", "bb"}, "--"))
}

func add(a, b int) {
	fmt.Printf("%d + %d = %d\n", a, b, a+b)
}
