package main

import "fmt"

func main() {
	var mes int8
	fmt.Println("Digite um mês(1 a 12): ")
	fmt.Scanf("%d", &mes)
	fmt.Printf("Estação nesse mês: %s", season(mes))
}

func season(month int8) string {
	var season string

	switch month {
	case 12, 1, 2:
		season = "Verão"
	case 3, 4, 5:
		season = "Outono"
	case 6, 7, 8:
		season = "Inverno"
	case 9, 10, 11:
		season = "Primavera"
	default:
		fmt.Println("Digite um mês válido!")
	}

	return season
}