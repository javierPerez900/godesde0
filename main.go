package main

import (
	"fmt"

	"github.com/javier/godesde0/variables"
)
func main() {
	// fmt.Println(variables.Nombre)
	// variables.RestoVariables()
	ok, texto := variables.ConviertoATexto(1567)
	fmt.Println(ok, texto)
	fmt.Printf("Tipo de texto: %T\n", texto)
}