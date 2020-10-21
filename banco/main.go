package main

import (
	"cursos/OOgo/BancoEmGo/banco/contas"
	"fmt"
)

//PagarBoleto método de pagar um boleto de uma conta
func PagarBoleto(conta verificarConta, valorDoBoleto float64) {
	conta.Sacar(valorDoBoleto)
}

type verificarConta interface {
	Sacar(valor float64) string
}

func main() {
	contaDoDenner := contas.ContaPoupanca{}
	contaDoDenner.Depositar(100)
	PagarBoleto(&contaDoDenner, 60)

	fmt.Println(contaDoDenner.ObterSaldo())

	contaDoTemer := contas.ContaCorrente{}
	contaDoTemer.Depositar(500)
	PagarBoleto(&contaDoTemer, -140)

	fmt.Println(contaDoTemer.ObterSaldo())
}
