package main

import "fmt"

func fibonacci() func() int {
	var x, y, sum = 0, 1, 0
	return func() int {
		sum = x + y
		x = y
		y = sum
		return sum
	}
}

func main() {
	f := fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Println(f())
	}
}