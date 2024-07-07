package main

import (
	"banco/contas"
	"fmt"
)

func PagarBoleto(conta verificarConta, valorDoBoleto float64) {
	conta.Sacar(valorDoBoleto)
}

type verificarConta interface {
	Sacar(valor float64) string
}

func main() {
	contaDoAndre := contas.ContaPoupanca{}
	contaDoAndre.Depositar(100)
	PagarBoleto(&contaDoAndre, 60)

	fmt.Println(contaDoAndre.ObterSaldo())

	contaDaCarol := contas.ContaCorrente{}
	contaDaCarol.Depositar(500)
	PagarBoleto(&contaDaCarol, 400)

	fmt.Println(contaDaCarol.ObterSaldo())

}
