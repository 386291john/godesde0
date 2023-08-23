package arreglos_slices

import "fmt"

var tabla [10]int = [10]int{10, 0, 0, 58, 69}
var matriz [5][10]int

func MuestroArreglo() {
	tabla[7] = 33
	tabla[2] = 54
	tabla2 := [10]string{"Pablo", "Pedro"}
	fmt.Println(tabla)
	fmt.Println(tabla2)

	for i := 0; i < len(tabla); i++ {
		fmt.Println(tabla[i])
	}
	matriz[0][5] = 15
	fmt.Println(matriz)
}
