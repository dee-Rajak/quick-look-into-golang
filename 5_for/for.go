package main

func main() {
	// i := 1

	//While loop
	// for i <= 5 {
	// 	// println(i)
	// 	// print("i: %v\t", i) i: %v   1i: %v  2i: %v  3i: %v  4i: %v  5
	// 	i = i + 1
	// }

	// infinite loop
	// for {
	// 	println("1")
	// }

	// general for loop
	// for i := 1; i <= 5; i++ {
	// 	println(i)
	// }

	for i := range 5 {
		println(i)
	}

	for _, i := range []int{10, 20, 30} {
		println(i)
	}
}
