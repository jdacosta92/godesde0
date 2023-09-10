package ejercicios

import (
	"fmt"

	"github.com/jdacosta92/godesde0/interfaces"
)

func HumanosRespirando(hu interfaces.Humano) {
	hu.Respirar()
	fmt.Printf("Soy un %s, y estoy respirando \n", hu.Sexo())
}
