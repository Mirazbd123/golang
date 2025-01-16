package main

import "fmt"

type payment struct{}

func (p payment) MakePayment(amount float32) {
	// razorPaymentGW := razorPay{} // for 1st pgw
	// razorPaymentGW.Pay(amount)

	stribePaymentGW := stribe{} // for 2nd pgw
	stribePaymentGW.Pay(amount)
}

// 1st payment gateway start
type razorPay struct{}

func (r razorPay) Pay(amount float32) {
	// logic for razor payment
	fmt.Println("Making patment using razorPay : ", amount)
}

// end

// 2nd Payment gateway start
type stribe struct{}

func (s stribe) Pay(amount float32) {
	fmt.Println("Making payment using stribe : ", amount)
}

//end

func main() {
	// razorPayment := payment{}
	// razorPayment.MakePayment(123)
	//when i use payment struct and a mathod of that struct Makepayment it will go to the MakePayment mathod and call the razorPay struct and
	// use its pay mathod and Print the line.
	// But it gets complicated when we add another payment gateway for payment like stribe. So that we have to make another struct for stribe
	// and make another mathod of that struct as Pay
	stribePayment := payment{}
	stribePayment.MakePayment(2344)
	// And for that we have to manually go to the MakePayment mathod and manualy change the calling stuct and its mathod and this is not a good
	// Practice and it will violate the open close principle which is mathods are open for extansion and close for modification

}
