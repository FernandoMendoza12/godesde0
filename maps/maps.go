package maps

import "fmt"

func MostrarMapa() {
	paises := make(map[string]string)

	paises["Mexico"] = "D.F."
	paises["Argentina"] = "Buenos Aires"
	fmt.Println(paises["Mexico"])

}
