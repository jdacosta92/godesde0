package files

import (
	//"bufio"
	"bufio"
	"fmt"

	//"ioutil"
	"os"
)

var fileName = "./files/txt/tabla.txt"

func GrabaTabla(tabla string) {
	archivo, err := os.Create(fileName)

	if err != nil {
		fmt.Println("Error al crear el archivo" + err.Error())
		return
	}

	fmt.Fprintln(archivo, tabla)
	archivo.Close()
}

func ReadFile() {
	archivo, err := os.ReadFile(fileName)
	if err != nil {
		fmt.Println("Error al leer archivo " + err.Error())
		return
	}

	fmt.Println(string(archivo))
}

func ReadFile2() {
	archivo, err := os.Open(fileName)
	if err != nil {
		fmt.Println("Error al leer archivo " + err.Error())
		return
	}

	scanner := bufio.NewScanner(archivo)
	for scanner.Scan() {
		registro := scanner.Text()
		fmt.Println("> " + registro)
	}
	archivo.Close()
}
