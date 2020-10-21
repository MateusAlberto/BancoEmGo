package contas

import "cursos/OOgo/BancoEmGo/banco/clientes"

//ContaPoupanca de um banco
//Deve ser iniciado com letra maiúscula para ter acesso público
//Deve ser iniciado com letra minúscula para ter acesso privado
type ContaPoupanca struct {
	Titular                              clientes.Titular
	NumeroAgencia, NumeroConta, Operacao int
	saldo                                float64
}

//Sacar método para realizar um saque de uma conta poupanca
func (c *ContaPoupanca) Sacar(valorDoSaque float64) string {
	podeSacar := valorDoSaque > 0 && valorDoSaque <= c.saldo
	if podeSacar {
		c.saldo -= valorDoSaque
		return "Saque realizado com sucesso"
	}
	return "saldo insuficiente"
}

//Depositar método para realizar o depósito de uma conta poupanca
func (c *ContaPoupanca) Depositar(valorDoDeposito float64) (string, float64) {
	if valorDoDeposito > 0 {
		c.saldo += valorDoDeposito
		return "Depósito realizado com sucesso", c.saldo
	}
	return "O valor do depósito deve ser maior que zero", c.saldo
}

//ObterSaldo funçao para retornar o valor do saldo
func (c *ContaPoupanca) ObterSaldo() float64 {
	return c.saldo
}
