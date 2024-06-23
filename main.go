package main

import "fmt"

type ContaCorrente struct {
	titular       string
	numeroAgencia int
	numeroConta   int
	saldo         float64
}

func main() {
	contaDoAndre := ContaCorrente{titular: "André",
		numeroAgencia: 589,
		numeroConta:   123456,
		saldo:         125.5}

	contaDaBruna := ContaCorrente{"Bruna", 222, 111222, 200}

	contaDaCarol := ContaCorrente{titular: "Carol", saldo: 125.5}
	fmt.Println(contaDoAndre)
	fmt.Println(contaDaBruna)
	fmt.Println(contaDaCarol)

	var contaDaCris *ContaCorrente
	contaDaCris = new(ContaCorrente)
	contaDaCris.titular = "Cris"
	contaDaCris.saldo = 500

	fmt.Println(*contaDaCris)

}
