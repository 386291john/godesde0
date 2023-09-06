package mapas

import "fmt"

func MostrarMapas() {
	paises := make(map[string]string)
	//fmt.Println(paises)

	paises["Mexico"] = "D.F."
	paises["Colombia"] = "Bogota"
	/*fmt.Println(paises)
	fmt.Println(paises["Colombia"])*/

	campeonato := map[string]int{
		"Barcelona":    39,
		"Nacional":     50,
		"Real Madrid":  45,
		"Boca Juniors": 15,
	}
	fmt.Println(campeonato)

	/*for equipo, puntaje := range campeonato {
		fmt.Printf("Equipo %s , tiene un puntaje de %d \n", equipo, puntaje)
	}*/
	delete(campeonato, "Real Madrid")
	fmt.Println(campeonato)

	puntaje, existe := campeonato["Nacional"]
	fmt.Printf("El puntaje capturado en %d, y el equipo existe = %t", puntaje, existe)
}
