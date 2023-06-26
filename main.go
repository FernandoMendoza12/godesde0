package main

import (
	ejerinterfaces "github.com/FernandoMendoza12/godesde0/ejer_Interfaces"
	"github.com/FernandoMendoza12/godesde0/modelos"
)

func main() {
	/*var a int = 10
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

	files.SumaTabla()
	files.LeerArchivo()

	funciones.Calculos()

	maps.MostrarMapa()

	structs.AltaUsuario()
	*/
	Pedro := new(modelos.Hombre)
	ejerinterfaces.HumanosRespirando(Pedro)

	Maria := new(modelos.Mujer)
	ejerinterfaces.HumanosRespirando(Maria)
}
