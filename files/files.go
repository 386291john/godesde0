package files

import (
	"bufio"
	"fmt"
	"github.com/386291john/godesde0/ejercicios"
	"io/ioutil"
	"os"
)

var filename string = "./files/txt/tabla.txt"

func GrabaTabla() {
	var texto string = ejercicios.Ingresos02()
	archivo, err := os.Create(filename)
	if err != nil {
		fmt.Println("Error al crear el archivo")
		return
	}
	fmt.Fprintln(archivo, texto)
	archivo.Close()
}

func SumaTabla() {
	var texto string = ejercicios.Ingresos02()
	if Append(filename, texto) == false {
		fmt.Println("error al concatenar contenido")
	}
}
func Append(filen string, texto string) bool {
	arch, err := os.OpenFile(filename, os.O_WRONLY|os.O_APPEND, 0644)
	if err != nil {
		fmt.Println("Error durante el Append" + err.Error())
		return false
	}
	_, err = arch.WriteString(texto)

	if err != nil {
		fmt.Println("Error durante el WriteString" + err.Error())
		return false
	}
	arch.Close()
	return true
}

func LeoArchivo() {
	archivo, err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Println("Error al leer archivo" + err.Error())
		return
	}
	fmt.Println(string(archivo))

}

func LeoArchivo2() {
	archivo, err := os.Open(filename)
	if err != nil {
		fmt.Println("Error al leer archivo" + err.Error())
		return
	}
	scanner := bufio.NewScanner(archivo)

	for scanner.Scan() {
		registro := scanner.Text()
		fmt.Println(">" + registro)

	}
	archivo.Close()
}
