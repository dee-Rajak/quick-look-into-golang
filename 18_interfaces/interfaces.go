package main

import "fmt"

type paymenter interface {
	pay(amount float32)
}

type payment struct {
	// gateway stripe
	gateway paymenter
}

// In solid, we voilating OpenClose Principle

func (p payment) makePayment(amount float32) {
	// razorpayPaymentGw := razorpay{}
	// razorpayPaymentGw.pay(amount)
	// stripePaymentGw := stripe{}
	p.gateway.pay(amount)
}

type razorpay struct{}

func (r razorpay) pay(amount float32) {
	// logic to make payment
	fmt.Println("making payment using razorpay", amount)
}

type stripe struct{}

func (s stripe) pay(amount float32){
	fmt.Println("making payment using stripe", amount)
}

type fakePayment struct {}

func (f fakePayment) pay(amount float32) {
	fmt.Println("making payment using fakePayment", amount)
}

func main() {
	// newPayment := payment{}
	// stripePaymentGw := stripe{}
	// razorpayPaymentGw := razorpay{}
	fakePaymentGw := fakePayment{}

	newPayment := payment{
		gateway: fakePaymentGw,
	}

	newPayment.makePayment(100.99)
}