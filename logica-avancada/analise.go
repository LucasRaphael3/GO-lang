package main

import "fmt"

func AnaliseNumeros() {
    var numeros [5]float64
    var soma float64
    var maior, menor float64

    fmt.Println("\n--- Análise de Números ---")
    for i := 0; i < 5; i++ {
        fmt.Printf("Digite o %dº número: ", i+1)
        fmt.Scanln(&numeros[i])
        soma += numeros[i]

        if i == 0 || numeros[i] > maior {
            maior = numeros[i]
        }
        if i == 0 || numeros[i] < menor {
            menor = numeros[i]
        }
    }

    media := soma / 5

    fmt.Println("Maior número:", maior)
    fmt.Println("Menor número:", menor)
    fmt.Printf("Média: %.2f\n", media)
}
