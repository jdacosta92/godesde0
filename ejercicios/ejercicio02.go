package ejercicios

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var texto string

func InputNum() (bool, int) {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("Ingreso un numero : ")
	if scanner.Scan() {
		input, err := strconv.Atoi(scanner.Text())
		if err != nil {
			panic("El dato ingresado no es un numero valido --> :" + err.Error())
		}
		return true, input
	}
	return false, 0
}

func TablaMultiplicacion(num int) string {
	for i := 0; i <= 10; i++ {
		//fmt.Printf("%d x %d = %d \n", num, i, num*i)
		texto += fmt.Sprintf("%d x %d = %d \n", num, i, num*i)
	}
	return texto
}
