package arreglosslices

import "fmt"

var tablaNumeros [10][10]int
var tablaNombres [10]string = [10]string{"Fernando", "Leia"}

func MostrarArreglos() {
	tablaNumeros[5][3] = 15
	fmt.Println(tablaNumeros)
	fmt.Println(tablaNombres)
}
