package main

import (
	"github.com/FernandoMendoza12/godesde0/middleware"
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

	Pedro := new(modelos.Hombre)
	ejerinterfaces.HumanosRespirando(Pedro)

	Maria := new(modelos.Mujer)
	ejerinterfaces.HumanosRespirando(Maria)

	deferPanic.EjemploPanic()

	ch1 := make(chan bool)
	go goroutines.MiNombreLento("Fernando Mendoza Moreno", ch1)
	defer func() {
		<-ch1
	}()
	fmt.Println("Estoy aqui")


	server.WebServer()*/
	middleware.MiMiddleware()

}
