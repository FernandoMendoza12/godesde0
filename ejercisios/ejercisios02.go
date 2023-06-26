package ejercisios

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func CapturarDatos() string {
	fmt.Println("Ingresa el numero que desea multiplicar :")
	var s string = ""
	scanner := bufio.NewScanner(os.Stdin)
	if scanner.Scan() {
		num, err := strconv.Atoi(scanner.Text())
		if err != nil {
			fmt.Println("El dato ingresado no es un numero", err)
		} else {
			fmt.Printf("Tabla del %d \n", num)
			for j := 0; j <= 10; j++ {
				s += fmt.Sprintf("%d x %d es igual a %d \n", num, j, num*j)
			}
		}
	}

	return s
}
