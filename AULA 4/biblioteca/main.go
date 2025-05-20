package main

import "fmt"

type Livro struct {
	Titulo  string
	Autor string
	AnoDeLançamento uint16
}
func (l Livro) printbonito(){
	fmt.Printf("Titulo: %s \nAutor: %s \nAno de Lançamento: %d", l.Titulo, l.Autor, l.AnoDeLançamento)
}

func main() {
	livro := Livro{Titulo: "A Morte de Ivan Ilitch", Autor: "Leo Tolstoy", AnoDeLançamento: 1886}
	livro.printbonito()
}