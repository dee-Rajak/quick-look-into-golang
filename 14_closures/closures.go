package main

import "fmt"

func counter() func() int {
	var count int = 0

	return func() int {
		count++
		return count
	}
}

func main() {
	// closures
	// anonymous function
	// function literal
	// function value
	// function closure
	// function that captures variables from outside its body
	// function that references variables from outside its body
	// function that is bound to its surrounding state
	// function that encloses its environment
	// function that captures the context in which it was created
	// function that captures the scope in which it was created

	increment := counter()
	increment2 := counter()

	fmt.Println(increment())
	fmt.Println(increment())
	fmt.Println(increment())
	fmt.Println(increment2())
	fmt.Println(increment2())
	fmt.Println(increment())
	fmt.Println(increment())
	fmt.Println(increment2())
}