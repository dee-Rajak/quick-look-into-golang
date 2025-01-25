package main

import (
	"fmt"
	"time"
)

type customer struct {
	id    int
	name  string
	phone string
}

type order struct {
	id        string
	amount    float32
	status    string
	createdAt time.Time
	customer
}

func main() {
	newOrder := order{
		id:        "1",
		amount:    10.99,
		status:    "received",
		createdAt: time.Now(),
	}

	fmt.Println("Order struct: ", newOrder)

	newOrder.customer = customer{
		id:    2,
		name:  "John",
		phone: "1234567890",
	}

	fmt.Println("Order struct: ", newOrder)
	fmt.Println("Customer name: ", newOrder.name)
	fmt.Println("Customer : ", newOrder.customer)
	// fmt.Println("ID: ", newOrder.id)
	// fmt.Println("Another Id: ", newOrder.customer.id)
	
	newOrder.customer.name = "Deez"
	fmt.Println("Customer : ", newOrder.customer)
}
