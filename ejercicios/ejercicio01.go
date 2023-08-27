package ejercicios

import (
	"strconv"
)

func ReturnIntString(texto string) (int, string) {
	condicion, _ := strconv.Atoi(texto)
	var resultado string

	if condicion > 100 {
		resultado = "Es mayor a 100"
	} else {
		resultado = "Es menor o igual a 100"
	}
	return condicion, resultado
}
