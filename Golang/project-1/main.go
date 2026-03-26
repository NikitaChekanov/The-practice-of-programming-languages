package main

import "fmt"

func main() {
	var nums = []int{}

	for i := 0; i < 10; i++ {
		nums = append(nums, i+1)
	}

	fmt.Println(nums)

	slice := make([]int, 1, 10)
	fmt.Println(slice)
}
