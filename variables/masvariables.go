package variables

import (
	"fmt"
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
