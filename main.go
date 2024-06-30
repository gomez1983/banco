package main

import (
	"banco/clientes"
	"banco/contas"
	"fmt"
)

func main() {
	clienteAndre := clientes.Titular{"André", "319.438.008-71", "Desenvolvedor Go"}
	contaDoAndre := contas.ContaCorrente{clienteAndre, 123, 123456, 100}
	// contaDoAndre := contas.ContaCorrente{Titular: clientes.Titular{
	// 	Nome:      "André",
	// 	CPF:       "319.438.008-71",
	// 	Profissao: "Desenvolvedor"},
	// 	NumeroAgencia: 123, NumeroConta: 123456, Saldo: 100}

	fmt.Println(contaDoAndre)
}
