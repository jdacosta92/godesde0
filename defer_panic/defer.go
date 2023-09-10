package defer_panic

import (
	"fmt"
	"log"
)

func VemosDefer() {
	// defer = instruccion que indica lo ultimo que se debe ejecutar al salir de la funcion
	fmt.Println("Este es un primer mensaje")

	defer fmt.Println("Este es el mensaje final")

	fmt.Println("Este es un segundo mensaje")
}

func EjemploPanic() {
	defer func() {
		reco := recover()
		if reco != nil {
			log.Fatalf("Ocurrio un error que genero un panic \n %v", reco)
		}
	}()
	a := 1
	if a == 1 {
		panic("Se encontro el valor 1")
	}
}
