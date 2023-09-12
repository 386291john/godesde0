package ejer_interfaces

import (
	"fmt"
	"github.com/386291john/godesde0/interfaces"
)

func HumanosRespirando(hu interfaces.Humano) {

	fmt.Printf("Soy un/a %s, y estoy respirando \n", hu.Sexo())
	fmt.Printf("Estoy comiendo? %t \n", hu.Comer())
	fmt.Printf("Estoy pensando? %t \n", hu.Pensar())
	fmt.Printf("Estoy respirando? %t \n \n", hu.Respirar())
	//fmt.Printf("Estoy comiendo? %t \n", hu.(*modelos.Hombre).Comiendo)
	//fmt.Printf("Estoy comiendo? %t \n", hu.(*modelos.Mujer).Comiendo)
}

func HumanosComiendo(hu interfaces.Humano) {

	fmt.Printf("Estoy comiendo? %t \n")
}
