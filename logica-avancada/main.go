package main

import (
    "fmt"
)

func main() {
    for {
        fmt.Println("\n--- Menu ---")
        fmt.Println("1. Tabuada Completa")
        fmt.Println("2. Cadastro com Repetição")
        fmt.Println("3. Análise de Números")
        fmt.Println("0. Sair")
        fmt.Print("Escolha uma opção: ")

        var opcao int
        fmt.Scanln(&opcao)

        switch opcao {
        case 1:
            Tabuada()
        case 2:
            Cadastro()
        case 3:
            AnaliseNumeros()
        case 0:
            fmt.Println("Encerrando programa.")
            return
        default:
            fmt.Println("Opção inválida!")
        }
    }
}
