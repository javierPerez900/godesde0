package variables

import (
	"fmt"
	"strconv"
	"time"
)

var Nombre string = "Javier"
var Estado bool
var Sueldo float32
var Fecha time.Time

func RestoVariables() {
	fmt.Println(Nombre)
	Nombre = "Juan"
	Estado = true
	Sueldo = 1234.56
	Fecha = time.Now()
	fmt.Println(Nombre)
	fmt.Println(Estado)
	fmt.Println(Sueldo)
	fmt.Println(Fecha)
}

func ConviertoATexto(numero int) (bool, string) {
	texto := strconv.Itoa(numero)
	return true, texto
}