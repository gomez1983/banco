package main

import (
	"banco/contas"
	"fmt"
)

func main() {
	contaDoAndre := contas.ContaPoupanca{}
	contaDoAndre.Depositar(100)
	contaDoAndre.Sacar(55)

	fmt.Println(contaDoAndre.ObterSaldo())

}
