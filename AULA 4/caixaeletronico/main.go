package main

import "fmt"

type Conta struct {
	Titular string
	Saldo float64
}
func (c *Conta) exibirExtrato(){
	fmt.Printf("Saldo da Conta: R$%2.f\n", c.Saldo)
}

func (c *Conta) depositaValor(valor float64) {
	if valor > 0{
		c.Saldo += valor
		fmt.Printf("Fazendo um deposito de R$%2.f\n", valor)
	} else{
		fmt.Println("---Valor invalido---")
	}
	c.exibirExtrato()
	
}

func (c *Conta) sacarValor(valor float64) {
	if valor > c.Saldo{
		fmt.Println("---Válor invalido---")
	} else{
		c.Saldo -= valor
		fmt.Printf("Tentando sacar R$%2.f da conta de: %s\n", valor, c.Titular)
	}
}

func (contaSaida Conta) TransferirValor(ContaEntrada Conta, valor float64){
	if valor <= contaSaida.Saldo {
		contaSaida.sacarValor(valor)
		ContaEntrada.depositaValor(valor)
	}
}

func main() {
	contaLucas := Conta{"Lucas", 0}
	contaLucas.depositaValor(50)

	contaPedro := Conta{"Pedro", 150}
	contaPedro.TransferirValor(contaLucas, 50)
	
}