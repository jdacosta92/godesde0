package arreglos_slices

import (
	"fmt"
)

var tabla [10]int = [10]int{10, 0, 59, 43} // indico que tabla es un array de 10 elementos
var matrix [10][10]int

func MuestroArreglos() {
	tabla[7] = 33
	tabla[2] = 54

	tabla2 := [10]string{"hola", "mundo"}
	fmt.Println(tabla2)

	for i := 0; i < len(tabla); i++ {
		fmt.Println(tabla[i])
	}
	fmt.Println("-----------------------------------")

	matrix[7][7] = 7
	fmt.Println(matrix)

	fmt.Println("-----------------------------------")
}
