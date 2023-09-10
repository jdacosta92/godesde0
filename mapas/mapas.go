package mapas

import (
	"fmt"
)

func MostrarMapas() {
	paises := make(map[string]string)
	//fmt.Println(paises)

	paises["mexico"] = "DF"
	paises["Argentina"] = "Buenos Aires"

	fmt.Println(paises)
	fmt.Println(paises["Argentina"])

	// si construye el mapa directamente con <map> me obliga a popularlo
	campeonato := map[string]int{
		"Barcelona":    39,
		"Real Madrid":  38,
		"Chivas":       37,
		"Boca Juniors": 30,
	}

	fmt.Println(campeonato)

	/*
		for equipo, puntaje := range campeonato {
			fmt.Printf("%s, tiene un puntaje de %d \n", equipo, puntaje)
		} */

	delete(campeonato, "Real Madrid")

	puntaje, existe := campeonato["Juventus"]
	fmt.Printf("El puntaje capturado es %d y el equipo existe : %t", puntaje, existe)

	fmt.Println()
	puntaje, existe = campeonato["Chivas"]
	fmt.Printf("El puntaje capturado es %d y el equipo existe : %t", puntaje, existe)

}
