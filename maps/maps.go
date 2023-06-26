package maps

import "fmt"

func MostrarMapa() {
	/*
		paises := make(map[string]string)

		paises["Mexico"] = "D.F."
		paises["Argentina"] = "Buenos Aires"
		fmt.Println(paises)
		fmt.Println(paises["Mexico"])
	*/

	campeonato := map[string]int{
		"Barcelona":   39,
		"Real Madrid": 38,
		"Chivas":      37,
		"Boca Jr":     30,
	}

	fmt.Println(campeonato)

	for k, v := range campeonato {
		fmt.Printf("El equipos %s, tiene un puntaje %d\n", k, v)
	}

	delete(campeonato, "Real Madrid")
	fmt.Println(campeonato)

	puntaje, existe := campeonato["Barcelona"]
	fmt.Printf("El puntaje capturado es de  : %d y el equipos es : %v \n", puntaje, existe)
}
