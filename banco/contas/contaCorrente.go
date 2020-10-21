package contas

import "cursos/OOgo/BancoEmGo/banco/clientes"

//ContaCorrente de um banco
//Deve ser iniciado com letra maiúscula para ter acesso público
//Deve ser iniciado com letra minúscula para ter acesso privado
type ContaCorrente struct {
	Titular                    clientes.Titular
	NumeroAgencia, NumeroConta int
	saldo                      float64
}

//Sacar método para realizar um saque de uma conta corrente
func (c *ContaCorrente) Sacar(valorDoSaque float64) string {
	podeSacar := valorDoSaque > 0 && valorDoSaque <= c.saldo
	if podeSacar {
		c.saldo -= valorDoSaque
		return "Saque realizado com sucesso"
	}
	return "saldo insuficiente"
}

//Depositar método para realizar o depósito de uma conta corrente
func (c *ContaCorrente) Depositar(valorDoDeposito float64) (string, float64) {
	if valorDoDeposito > 0 {
		c.saldo += valorDoDeposito
		return "Depósito realizado com sucesso", c.saldo
	}
	return "O valor do depósito deve ser maior que zero", c.saldo
}

//Transferir método para realizar uma transferência
//entre esta conta corrente e a passada como parâmetro
func (c *ContaCorrente) Transferir(valorTransferencia float64, contaDestino *ContaCorrente) bool {
	if valorTransferencia > 0 && valorTransferencia <= c.saldo {
		c.saldo -= valorTransferencia
		contaDestino.Depositar(valorTransferencia)
		return true
	}
	return false
}

//ObterSaldo funçao para retornar o valor do saldo
func (c *ContaCorrente) ObterSaldo() float64 {
	return c.saldo
}
