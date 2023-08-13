package ejercicios

import (
	"strconv"
)

var texto string

func Ejercicio(valor string) (int, string) {
	num, _ := strconv.Atoi(valor)

	if num > 100 {
		texto = "Es mayor a 100"
		//return (num, "Es mayor a 100")
		//fmt.Println(num, "Es mayor a 100")
	} else {
		texto = "Es menor a 100"
		//return (num, "Es menor a 100")
		//fmt.Println(num, "Es menor a 100")
	}
	return num, texto
}
