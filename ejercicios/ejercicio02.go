package ejercicios

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var numero1 int
var err error

func Ingresos02() {
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
		println(numero1, "*", i, "=", numero1*i)
	}

}
