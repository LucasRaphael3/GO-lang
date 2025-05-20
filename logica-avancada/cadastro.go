package main

import (
    "bufio"
    "fmt"
    "os"
    "strings"
)

func Cadastro() {
    var nomes []string
    scanner := bufio.NewScanner(os.Stdin)

    fmt.Println("\n--- Cadastro de Nomes ---")
    for {
        fmt.Print("Digite um nome (ou 'sair' para encerrar): ")
        scanner.Scan()
        nome := scanner.Text()
        if strings.ToLower(nome) == "sair" {
            break
        }
        nomes = append(nomes, nome)
    }

    fmt.Println("\nNomes cadastrados:")
    for _, nome := range nomes {
        fmt.Println(nome)
    }
}
