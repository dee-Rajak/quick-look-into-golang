package main

import "fmt"

func main() {
	// numbered sequence of specific length
	
	var nums [5]int
	nums[0] = 5
	nums[2] = 9
	fmt.Println(len(nums), cap(nums))
	fmt.Println(nums)
}