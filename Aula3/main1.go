package main

import "fmt"

func main() {
	var soma float32
	var notas = [5]float32{7.5, 6.5, 4, 8.5, 10}
	for i, v := range notas{
		fmt.Printf("Nota %d: %2.f\n", i+1, v)
		soma += v
	}
	media := soma/5
	fmt.Printf("Media: %f", media)

	
}