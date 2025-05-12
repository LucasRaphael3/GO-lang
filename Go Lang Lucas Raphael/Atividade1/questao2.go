package main

import "fmt"

func main() {
	var idade int8
	fmt.Println("Digite sua idade: ")
	fmt.Scanf("%d", &idade)
	fmt.Printf("Sua faixa etária: %s", age(idade))
	
}

func age(idade int8) string {
	var faixaeta string

	switch {
	case idade < 12:
		faixaeta = "Criança"
	case idade < 18:
		faixaeta = "Adolescente"
	case idade < 60:
		faixaeta = "Adulto"
	case idade > 59:
		faixaeta = "Idoso"
	default:
		fmt.Println("Digite uma idade válida!")
	}

	return faixaeta
}
