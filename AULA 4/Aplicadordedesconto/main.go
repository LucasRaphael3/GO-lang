package main

import "fmt"

type Produto struct {
	Nome string
	Preço float64
}

func (p *Produto) AplicarDesconto(valor float64) {
	if valor > 0{
		p.Preço = p.Preço - ((p.Preço * valor)/100)
		fmt.Printf("\nNovo valor do %s com o desconto de %.f porcento é de R$%2.f", p.Nome, valor, p.Preço)
	} else{
		fmt.Println("Valor Invalido")
	}
}

func (p *Produto) ExibirProduto() {
	fmt.Printf("\nNome: %s \nValor: R$%2.f\n", p.Nome, p.Preço)
}
func main() {
	arroz := Produto{"Arroz", 20}
	arroz.ExibirProduto()
	arroz.AplicarDesconto(25)
	arroz.ExibirProduto()
	feijao := Produto{"Feijao", 30}
	feijao.ExibirProduto()
	feijao.AplicarDesconto(15)
	feijao.ExibirProduto()
	cafe := Produto{"Cafe", 35}
	cafe.ExibirProduto()
	cafe.AplicarDesconto(5)
	cafe.ExibirProduto()
}