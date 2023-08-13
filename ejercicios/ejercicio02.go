package ejercicios

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var numero1 int
var err error
var texto string

func Ingresos02() string {
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Println("Ingrese numero : ")
		if scanner.Scan() {
			numero1, err = strconv.Atoi(scanner.Text())
			if err != nil {
				continue
			} else {
				break
			}
		}
	}
	for i := 1; i < 10; i++ {
		//fmt.Println(numero1, "*", i, "=", numero1*i)
		texto += fmt.Sprintln(numero1, "*", i, "=", numero1*i)

	}
	return texto
}
