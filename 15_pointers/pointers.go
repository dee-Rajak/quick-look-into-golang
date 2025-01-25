package main

import "fmt"

// Pass by value
// func changeNum(num int) {
// 	num = 5
// 	fmt.Println("In changeNum:", num)
// 	// fmt.Println("In changeNum:", *num)
// 	fmt.Println("In changeNum, num var mem address", &num)
// }

// Pass by reference
func changeNum(num *int) {
	*num = 5
	fmt.Println("In changeNum:", num)
	fmt.Println("In changeNum:", *num)
	// fmt.Println("In changeNum:", *num)
	fmt.Println("In changeNum, num var mem address", &num)
}

func main() {
	num := 10

	fmt.Println("Before changeNum in main:", num)
	fmt.Println("Before changeNum in main, num var mem address", &num)

	changeNum(&num)

	fmt.Println("After changeNum in main:", num)
	// fmt.Println("After changeNum in main:", *num)
	fmt.Println("After changeNum in main, num var mem address", &num)
}