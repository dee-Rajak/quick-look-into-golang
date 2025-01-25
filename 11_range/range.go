package main

import "fmt"

func main() {
	// nums := []int{2, 3, 4}

	// for i := 0; i < len(nums); i++ {
	// 	println(nums[i])
	// }

	// for i, v := range nums {
	// 	println(i, v)
	// }

	// sum := 0
	// for _, v := range nums {
	// 	sum += v
	// }

	// println("sum:", sum)

	// m := map[string]string{
	// 	"fname": "John",
	// 	"lname": "Doe",
	// }

	// for k, v := range m {
	// 	println(k, v)
	// }

	// unicode code point rune
	// starting byte of rune
	for i, c := range "golang" {
		// fmt.Println(i, c)
		fmt.Println(i, string(c))
	}
}
