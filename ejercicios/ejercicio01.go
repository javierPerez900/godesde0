package ejercicios

import "strconv"

func DevuelveValores(cadena string) (int, string) {
	var mensaje string
	entero, _ := strconv.Atoi(cadena)
	if entero > 100 {
		mensaje = "Es mayor que 100"
	} else {
		mensaje = "Es menor  100"
	}
	return entero, mensaje
}