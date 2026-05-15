package main

import (
	"fmt"
	"runtime"

	"github.com/javier/godesde0/ejercicios"
	"github.com/javier/godesde0/variables"
)
func main() {
	fmt.Println(variables.Nombre)
	// variables.RestoVariables()
	// ok, texto := variables.ConviertoATexto(1567)
	// fmt.Println(ok, texto)
	// fmt.Printf("Tipo de texto: %T\n", texto)

	if os := runtime.GOOS; os == "Linux." || os == "OS X." {
		fmt.Println("Esto no es Windows")
	} else {
		fmt.Println("Esto es Windows")
	}

	switch os := runtime.GOOS; os {
	case "linux":
		fmt.Println("Esto es Linux")
	case "darwin":
		fmt.Println("Esto es Darwin")
	default:
		fmt.Printf("%s.\n", os)
	}
	entero, mensaje := ejercicios.DevuelveValores("200")
	fmt.Println(entero, mensaje)
}