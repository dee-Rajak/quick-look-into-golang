package main

import "fmt"

// struct is a composite data type
// struct is a collection of fields
// struct is a user-defined data type
// struct can be used for object-oriented programming
// we create methods for structs to define behavior
// use pointer receivers for methods that modify the value
// use value receivers for methods that do not modify the value

// type order struct {
// 	id        string
// 	amount    float32
// 	status    string
// 	createdAt time.Time // nanoseconds precision
// }

// func newOrder(id string, amount float32, status string) *order {
// 	ord := order{
// 		id: id,
// 		amount: amount,
// 		status: status,
// 	}
// 	return &ord
// }

// func (o *order) changeStatus(status string) {
// 	o.status = status
// }

// func (o order) getAmount() float32 {
// 	return o.amount
// }

func main() {
	// ord := order{
	// 	id: "1",
	// 	amount: 10.99,
	// 	status: "received",
	// }

	// ord.createdAt = time.Now()
	// fmt.Println("Order struct: ", ord)
	// fmt.Println("Order ID: ", ord.id)
	// fmt.Println("Order Amount: ", ord.amount)
	// fmt.Println("Order Status: ", ord.status)
	// fmt.Println("Order Created At: ", ord.createdAt.UTC())

	// ord := order{
	// 	id: "1",
	// 	amount: 10.99,
	// 	status: "received",
	// }

	// ord.changeStatus("shipped")
	// fmt.Println("Order struct: ", ord)

	// fmt.Println("Order Amount: ", ord.getAmount())

	// myOrd := newOrder("1", 10.99, "received")
	// fmt.Println("Order struct: ", myOrd)

	lang := struct {
		name string
		isGood bool
	} {"Go", true}

	fmt.Println(lang)
}
