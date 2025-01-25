package main

import "fmt"

// func printSlice(items []int) {
// 	for _, item := range items {
// 		fmt.Println(item)
// 	}
// }

// func printStringSlice(items []string) {
// 	for _, item := range items {
// 		fmt.Println(item)
// 	}
// }

func printSlice[T int | string](items []T) {
	for _, item := range items {
		fmt.Println(item)
	}
}

func main() {
	// nums := []int{1, 2, 3, 4, 5}
	// names := []string{"John", "Doe", "Elaine", "Deez"}
	// printSlice(nums)
	// printStringSlice(names)
	nums := []int{1, 2, 3, 4, 5}
	names := []string{"John", "Doe", "Elaine", "Deez"}
	printSlice(nums)
	printSlice(names)
}