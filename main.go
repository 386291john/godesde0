package main

import (
	"fmt"
	"github.com/386291john/godesde0/ejercicios"
)

func main() {
	/*estado, texto := variables.ConviertoaTexto(1588)
	fmt.Println(estado)
	fmt.Println(texto)*/

	num, texto := ejercicios.Ejercicio("dddd")
	fmt.Println(num)
	fmt.Println(texto)

	/*if os := runtime.GOOS; os == "linux" || os == "OS X." {
		fmt.Println("Esto no es Windows, es ", os)
	} else {
		fmt.Println("Esto es Windows")
	}

	switch os := runtime.GOOS; os {
	case "linux":
		fmt.Println("Esto es Linux")
	case "darwin":
		fmt.Println("esto es Darwin")
	default:
		fmt.Printf("%s \n", os)
	}*/
}
