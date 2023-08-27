package main

import (
	"fmt"

	"github.com/jdacosta92/godesde0/variables"
)

func main() {
	estado, resultado := variables.ConviertoaTexto(1588)
	fmt.Println(estado)
	fmt.Println(resultado)
}
