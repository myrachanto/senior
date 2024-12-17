package main

import "fmt"

type PaymentInterface interface {
	Payment() error
}

type Paypal struct {
	Name   string
	Email  string
	Amount float64
}

func NewPaypalPayment(name string, email string, amount float64) PaymentInterface {
	return &Paypal{
		Name:   name,
		Email:  email,
		Amount: amount,
	}
}
func (p *Paypal) Payment() error {
	//Todo implement payment with paypal
	fmt.Printf("processing payment of %.2f with paypal to %s with an email %s \n", p.Amount, p.Name, p.Email)
	return nil
}

type Stripe struct {
	Name   string
	Email  string
	Amount float64
}

func NewStripePayment(name string, email string, amount float64) PaymentInterface {
	return &Stripe{
		Name:   name,
		Email:  email,
		Amount: amount,
	}
}
func (p *Stripe) Payment() error {
	//Todo implement payment with stripe
	fmt.Printf("processing payment of %.2f with stripe to %s with an email %s \n", p.Amount, p.Name, p.Email)
	return nil
}

// Method Injection
func ProcessPayment(payment PaymentInterface) {
	err := payment.Payment()
	if err != nil {
		fmt.Println("Error processing payment:", err)
	}
}

// Field Injection
type PaymentProcessor struct {
	Payment PaymentInterface
}

func (pp *PaymentProcessor) Process() {
	if pp.Payment != nil {
		err := pp.Payment.Payment()
		if err != nil {
			fmt.Println("Error processing payment:", err)
		}
	} else {
		fmt.Println("No payment method provided!")
	}
}
func main() {
	// Constructor Injection
	paypalPayment := NewPaypalPayment("Alex Smith", "alexsmith@gmail.com", 4500)
	paypalPayment.Payment()
	stripePayment := NewStripePayment("Jennifer Gutierez", "jennyGutierez56@hotmail.com", 5000)
	stripePayment.Payment()

	// Method Injection
	ProcessPayment(paypalPayment)
	ProcessPayment(stripePayment)

	// Field Injection
	paypalProcessor := &PaymentProcessor{Payment: paypalPayment}
	paypalProcessor.Process()

	stripeProcessor := &PaymentProcessor{Payment: stripePayment}
	stripeProcessor.Process()
}
