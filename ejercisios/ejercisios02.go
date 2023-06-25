package ejercisios

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func CapturarDatos() {
	fmt.Println("Ingresa el numero que desea multiplicar :")
	scanner := bufio.NewScanner(os.Stdin)
	if scanner.Scan() {
		num, err := strconv.Atoi(scanner.Text())
		if err != nil {
			fmt.Println("El dato ingresado no es un numero", err)
		} else {
			mostrarTabla(num)
		}
	}
}

func mostrarTabla(i int) {
	for j := 0; j <= 10; j++ {
		fmt.Println(j * i)
	}
}
