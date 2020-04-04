package main

import "fmt"

type PayMethod interface {
	Pay()
}

type Paypal struct{}

func (p Paypal) Pay() {
	fmt.Println("Pagado por paypal")
}

type Cash struct{}

func (c Cash) Pay() {
	fmt.Println("Pagado por efectivo")
}

type CreditCard struct{}

func (c CreditCard) Pay() {
	fmt.Println("Pagado por tarjeta de crédito")
}

func Factory(method uint) PayMethod {
	switch method {
	case 1:
		return Paypal{}
	case 2:
		return Cash{}
	case 3:
		return CreditCard{}
	default:
		return nil
	}
}

func main() {
	var method uint
	fmt.Println("Digite un número de método de pago")
	fmt.Println("\t 1:Paypal 2:Efectivo 3:Tarjeta de crédito")
	_, err := fmt.Scanln(&method)
	if err != nil {
		panic("debe digitar un método valido")
	}

	if method > 3 {
		panic("debe digitar un método valido")
	}

	payMethod := Factory(method)
	payMethod.Pay()
}
