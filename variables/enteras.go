package variables

import (
	"fmt"
)

func MuestroEnteros() { // si el nombre de la func comienza con Mayusculas entonces se puede invocar desde fuera (desde el main por ej)
	var intcomun int      // declaracion de variable utilizando la palabra reservada "var"
	intde32 := int32(10)  // asignacion sin declaracion, asigno el valor 10 en formato int32bits
	intde64 := int64(100) // asignacion sin declaracion, asigno el valor 100 en formato int64bits

	fmt.Println("intcomun =", intcomun)
	fmt.Println("intde32 =", intde32)
	fmt.Println("intde64 =", intde64)
}
