package main

import (
	"fmt"
	"runtime"

	"github.com/FernandoMendoza12/godesde0/ejercisios"
	"github.com/FernandoMendoza12/godesde0/variables"
)

func main() {
	var a int = 10
	variables.MostrarEnteros(a)
	variables.MasVariables()
	fmt.Println(variables.ImprimirBoolYString())
	fmt.Println(variables.ConvertirTexto(a))

	os := runtime.GOOS

	if os == "Windows." {
		fmt.Println("Este sistemas es Linux")
	} else {
		fmt.Println("Este sistema no es linux")
	}

	value, text := ejercisios.RetornarDosValores("1750")

	fmt.Println(value)
	fmt.Println(text)
}
