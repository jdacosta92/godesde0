package main

import (
	"github.com/jdacosta92/godesde0/middleware"
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
	}

	files.ReadFile()

	funciones.Calculos()

	funciones.LlamarClosure() */

	//mapas.MostrarMapas()

	//users.AltaUsuario()

	//Pedro := new(modelos.Hombre)
	//ejercicios.HumanosRespirando(Pedro)

	//defer_panic.VemosDefer()
	//defer_panic.EjemploPanic()

	/*canal1 := make(chan bool)
	go goroutines.MiNombreLento("Julian", canal1)
	defer func() { <-canal1 }() // defer con todos los channel await

	fmt.Println("Estoy aqui")

	var x string
	fmt.Scanln(&x)
	fmt.Println("El usuario ingreso por teclado : ", x) */

	//<-canal1 // esto signica que va a esperar a que canal1 existe como verdadero (true)
	// se puede poner todos los channel al final o .. se puede poner todos en un defer arriba para mas prolijidad

	//webserver.MiWebServer()

	middleware.MiMiddleware()

}
