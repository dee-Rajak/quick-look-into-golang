package main

import "fmt"

func main() {
	// numbered sequence of specific length
	// slices are more flexible than arrays
	// slices are more common than arrays
	// slices are a reference to an array
	// slices are like dynamic arrays with no fixed size
	// slices can grow and shrink
	// slices are more powerful than arrays
	// slices are more efficient than arrays

	// uninitialized slice is nil
	// var nums []int
	// fmt.Println(nums == nil)
	// fmt.Println(nums, len(nums), cap(nums))
	
	var nums = make([]int, 5)
	fmt.Println(nums == nil)
	fmt.Println(nums, len(nums), cap(nums))
}