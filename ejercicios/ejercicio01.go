package ejercicios

import "strconv"

func DevuelveValores(cadena string) (int, string) {
	entero, err := strconv.Atoi(cadena)
	if err != nil {
		return 0, "Error al convertir la cadena a entero " + err.Error()
	}
	if entero > 100 {
		return entero, "Es mayor que 100"
	} else {
		return entero, "Es menor  100"
	}
}