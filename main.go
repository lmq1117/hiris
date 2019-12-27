package main

import (
	"fmt"
	"time"
)

func main() {
	for i := 0; i < 10; i++ {
		add(1, i)
	}
	time.Sleep(1e9)
}

func add(a, b int) {
	time.Sleep(3e9)
	c := a + b
	fmt.Printf("%d + %d = %d\n", a, b, c)

}
