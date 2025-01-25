package main

import "fmt"

func main() {

	m := make(map[string]string)

	m["name"] = "John"

	fmt.Println(m)
	fmt.Println(m["name"])
}