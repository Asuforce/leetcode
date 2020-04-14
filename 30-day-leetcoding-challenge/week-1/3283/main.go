package main

import "fmt"

func main() {
	i := []int{2, 2, 1}
	fmt.Println(singleNumber(i))
}

func singleNumber(nums []int) int {
	u := 0
	for _, n := range nums {
		cnt := 0
		for _, e := range nums {
			if n == e {
				cnt++
			}
		}
		if cnt == 1 {
			u = n
		}
	}
	return u
}
