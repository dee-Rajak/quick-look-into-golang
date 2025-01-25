package main

import "fmt"

func sum(a ...int) int {
	total := 0

	for _, num := range a {
		total += num
	}

	return total
}

func main() {
	// variadic function
	// accepts variable number of arguments
	// fmt.Println("Hello", "World")
	// fmt.Println("Hello", "World", "Go")
	// fmt.Println("Hello", "World", "Go", "Programming")

	// deez("Hello", "World", 5, true, 5.69)

	// fmt.Println("Hello", "World", 5, true, 5.69)

	
	result := sum(2, 3, 4, 5, 6, 7, 8, 9, 10)
	fmt.Println(result)
	fmt.Println(sum(2, 3, 4, 5, 6, 7, 8, 9, 10))
	
	nums := []int{2, 3, 4, 5, 6, 7, 8, 9, 10}
	fmt.Println(sum(nums...))	
}