package main

import "fmt"

func main() {
	fmt.Println(fib(4))
}

func fib(N int) int {
	switch N {
	case 0:
		return 0
	case 1:
		return 1
	default:
		return fib(N-1) + fib(N-2)
	}
}
