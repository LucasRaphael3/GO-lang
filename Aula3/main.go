package main

import "fmt"

func main() {
	quest5()
}

func quest1() {
	for i:= 2; i<=100; i++ {
		if i % 2 != 0{
			continue
		}
		fmt.Printf("%d - ", i)
	}
}

func quest2() {
	soma := 0
	for i:= 1; i<100; i++ {
		if i % 2 == 0{
			continue
		}
		soma += i
	}
	fmt.Printf("Soma de todos os impares de 1 a 100: %d", soma)
}

func quest3() {
	for i:= 0; i<=100; i++{
			senha := ""
			fmt.Println("Digite sua senha: ")
			fmt.Scanf("%s\n", &senha)
			if senha == "1234"{
				break
			}
		}
	fmt.Println("Acesso Liberado")
}

func quest4() {
	for i := 50; i > 0; i--{
			fmt.Printf("%d - ", i)
		}
}

func quest5() {
	for i:= 10; i<=25; i++ {
			for j:=1; j<=10; j++ {
				fmt.Printf("%d x %d = %d\n", i, j, i*j)
			}
		}
}