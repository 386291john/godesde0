package middlewares

import "fmt"

func sumar(a, b int) int {
	return a + b
}

func multiplicar(a, b int) int {
	return a * b
}

func restar(a, b int) int {
	return a - b
}

func MiMiddlewere() {
	fmt.Println("Inicio")

	result := operaconesMidd(sumar)(2, 3)
	fmt.Println(result)
	result = operaconesMidd(restar)(10, 6)
	fmt.Println(result)
	result = operaconesMidd(multiplicar)(2, 4)
	fmt.Println(result)
}
func operaconesMidd(f func(int, int) int) func(int, int) int {
	return func(a, b int) int {
		fmt.Println("Inicio de la operacion")
		return f(a, b)
	}
}
