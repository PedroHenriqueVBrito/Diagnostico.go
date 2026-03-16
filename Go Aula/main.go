package main

import "fmt"

func deposito(saldo *float64, valor float64) {
	*saldo += valor
}

func saque(saldo *float64, valor float64) {
	if valor > *saldo {
		fmt.Println("Saldo insuficiente")
		return
	}
	*saldo -= valor
}

func main() {

	saldo := 1000.0

	fmt.Printf("Saldo inicial: %.2f\n", saldo)

	deposito(&saldo, 200)
	fmt.Printf("Após depósito: %.2f\n", saldo)

	saque(&saldo, 150)
	fmt.Printf("Após saque: %.2f\n", saldo)

	saque(&saldo, 1200)
	fmt.Printf("Saldo da conta após saque insuficiente: %.2f\n", saldo)
}
