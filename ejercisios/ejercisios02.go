package ejercisios

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func capturarDatos() {
	fmt.Println("Ingresa el numero que desea multiplicar :")
	scanner := bufio.NewScanner(os.Stdin)
	if scanner.Scan() {
		num, err := strconv.Atoi(scanner.Text())
		if err != nil {
			mostrarTabla(num)
		} else {
			fmt.Println("El dato ingresado no es un numero")
		}
	}
}

func mostrarTabla(i int) {
	for j := 0; i < 10; i++ {
		fmt.Println(i * j)
	}
}
