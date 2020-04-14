package main

import "fmt"

func main() {
	printReverse(0, []int{1, 2, 3, 4, 5})
}

func printReverse(i int, nums []int) {
	if len(nums) == 0 || i >= len(nums) {
		return
	}
	printReverse(i+1, nums)
	fmt.Println(nums[i])
}
