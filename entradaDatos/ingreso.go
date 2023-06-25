package entradadatos

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var num1 int
var num2 int
var leyenda string
var err error

func IngresoNum() {
	fmt.Println("Ingrese el numero uno :")
	scanner := bufio.NewScanner(os.Stdin)
	if scanner.Scan() {
		num1, err = strconv.Atoi(scanner.Text())
		if err != nil {
			panic("No se ingreso ningun dato fue incorrecto" + err.Error())
		}
	}

	fmt.Println("Ingrese el numero dos :")
	if scanner.Scan() {
		num2, err = strconv.Atoi(scanner.Text())
		if err != nil {
			panic("No se ingreso ningun dato fue incorrecto" + err.Error())
		}
	}

	fmt.Println("Ingrese a leyenda :")
	if scanner.Scan() {
		leyenda = scanner.Text()
	}

	fmt.Println(leyenda, num1*num2)
}
