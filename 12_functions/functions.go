package main

// func add(x int, y int) int {
// 	return x + y
// }

func add(x, y int) int {
	return x + y
}

func getLanguages() (string, string, string) {
	return "Go", "Python", "JavaScript"
}

func processIt(fn func(a int) int) {
	fn(10)
}

func main() {
	// named sequence of statements that performs a computation
	// functions are reusable
	// functions are a way to divide the code into smaller parts
	// functions are a way to organize the code
	// functions are a way to make the code more readable
	// functions are a way to make the code more maintainable
	// functions are a way to make the code more testable
	// functions are a way to make the code more efficient
	// functions are a way to make the code more reusable
	// functions are a way to avoid code duplication
	// functions are a way to avoid bugs
	// functions are a way to avoid side effects
	// functions are a way to avoid spaghetti code
	// functions are a way to avoid callback hell

	// result := add(5, 9)
	// fmt.Println(result)

	// fmt.Println(getLanguages())

	// lang1, lang2, _ := getLanguages()
	// fmt.Println(lang1, lang2)

	// fn := func(a int) int {
	// 	return a
	// }

	// processIt(fn)
}
