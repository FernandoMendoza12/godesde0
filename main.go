package main

import (
	"fmt"

	"github.com/FernandoMendoza12/godesde0/variables"
)

func main() {
	var a int = 10
	variables.MostrarEnteros(a)
	variables.MasVariables()
	fmt.Println(variables.ImprimirBoolYString())
	fmt.Println(variables.ConvertirTexto(a))
}
