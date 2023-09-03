package main

import (
	"github.com/jdacosta92/godesde0/files"
)

func main() {
	/*estado, resultado := variables.ConviertoaTexto(1588)
	fmt.Println(estado)
	fmt.Println(resultado)

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

	result, msg := ejercicios.ReturnIntString("g")
	fmt.Println("Resultado :", result, "-->", msg)

	keyboard.IngresoNumeros()

	status, num := ejercicios.InputNum()

	if status {
		tabla := ejercicios.TablaMultiplicacion(num)
		files.GrabaTabla(tabla)
	} */

	files.ReadFile()
}
