
package main

import "fmt"

type Aluno struct {
	Nome  string
	Notas []float64
}

// Método para calcular média e imprimir resultado
func (a Aluno) CalcularMedia() {
	if len(a.Notas) == 0 {
		fmt.Printf("Aluno %s não possui notas.\n", a.Nome)
		return
	}

	var soma float64
	for _, nota := range a.Notas {
		soma += nota
	}
	media := soma / float64(len(a.Notas))

	status := "Reprovado"
	if media >= 6.0 {
		status = "Aprovado"
	}

	fmt.Printf("Aluno: %s | Média: %.2f | Situação: %s\n", a.Nome, media, status)
}

func main() {
	aluno1 := Aluno{
		Nome:  "Ana",
		Notas: []float64{7.5, 8.0, 6.5},
	}

	aluno2 := Aluno{
		Nome:  "Bruno",
		Notas: []float64{5.0, 4.5},
	}

	aluno3 := Aluno{
		Nome:  "Carla",
		Notas: []float64{10.0, 9.5, 9.0, 8.5},
	}

	aluno1.CalcularMedia()
	aluno2.CalcularMedia()
	aluno3.CalcularMedia()
}
