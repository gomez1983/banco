package main

import "fmt"

type ContaCorrente struct {
	titular       string
	numeroAgencia int
	numeroConta   int
	saldo         float64
}

func (c *ContaCorrente) Sacar(valorDoSaque float64) string {
	podeSacar := valorDoSaque > 0 && valorDoSaque <= c.saldo
	if podeSacar {
		c.saldo -= valorDoSaque
		return "Saque realizado com sucesso"
	} else {
		return "Saldo insuficiente"
	}
}

func main() {
	contaDoAndre := ContaCorrente{}
	contaDoAndre.titular = "André"
	contaDoAndre.saldo = 500

	fmt.Println(contaDoAndre.saldo)

	fmt.Println(contaDoAndre.Sacar(200))
	fmt.Println(contaDoAndre.saldo)
}
