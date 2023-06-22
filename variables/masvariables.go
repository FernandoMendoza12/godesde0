package variables

import (
	"fmt"
	"strconv"
	"time"
)

var name string = "Fernando"
var vivo bool = true
var sueldo float32 = 44000
var fecha time.Time

func MasVariables() {
	fmt.Println(name)
	fmt.Println(vivo)
	fmt.Println(sueldo)
	fmt.Println(fecha)
}

func ImprimirBoolYString() string {
	return fmt.Sprintf("This si %s and he is %v", name, vivo)
}

func ConvertirTexto(numero int) string {
	texto := strconv.Itoa(numero)
	return texto
}
