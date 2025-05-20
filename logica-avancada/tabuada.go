package main

import "fmt"

func Tabuada() {
    fmt.Println("\n--- Tabuada de 1 a 10 ---")
    for i := 1; i <= 10; i++ {
        for j := 1; j <= 10; j++ {
            fmt.Printf("%d x %d = %d\t", i, j, i*j)
        }
        fmt.Println()
    }
}
