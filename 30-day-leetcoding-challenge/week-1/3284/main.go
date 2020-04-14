package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(isHappy(35))
}

func isHappy(n int) bool {
	for {
		n = powNum(n)

		if n == 1 {
			return true
		} else if n < 5 {
			return false
		}
	}
}

func powNum(n int) (r int) {
	s := countDigits(n)
	for _, n := range s {
		r += int(math.Pow(float64(n), 2))
	}
	return
}

func countDigits(i int) (s []int) {
	for i != 0 {
		j := i
		i /= 10
		s = append(s, j%10)
	}
	return
}
