package main

import "fmt"

func main() {
	fmt.Println(getRow(4))
}

func getRow(rowIndex int) []int {
	r := make([]int, rowIndex+1)
	r[0] = 1

	for i := 0; i < rowIndex+1; i++ {
		for j := i; j >= 1; j-- {
			r[j] += r[j-1]
		}
	}
	return r
}
