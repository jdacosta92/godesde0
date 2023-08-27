package main

import (
	"fmt"
	"runtime"

	"github.com/jdacosta92/godesde0/ejercicios"
)

func main() {
	/*estado, resultado := variables.ConviertoaTexto(1588)
	fmt.Println(estado)
	fmt.Println(resultado)*/

	if os := runtime.GOOS; os == "linux" || os == "OS X." {
		fmt.Println("Esto no es windows", os)
	} else {
		fmt.Println("Esto es windows")
	}

	switch os := runtime.GOOS; os {
	case "linux":
		fmt.Println("Esto es Linux")
	case "darwin":
		fmt.Println("Esto es Darwin")
	default:
		fmt.Printf("%s \n", os)
	}

	result, msg := ejercicios.ReturnIntString("23")
	fmt.Println("Resultado :", result, "-->", msg)

}
