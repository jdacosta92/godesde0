package ejercicios

import (
	"strconv"
)

func ReturnIntString(texto string) (int, string) {
	condicion, err := strconv.Atoi(texto) // hay funciones que devuelven mas de un resultado
	var resultado string

	if err != nil { // manejo de errores?
		return 0, "Hubo un error : " + err.Error()
	}

	if condicion > 100 {
		resultado = "Es mayor a 100"
	} else {
		resultado = "Es menor o igual a 100"
	}
	return condicion, resultado
}
