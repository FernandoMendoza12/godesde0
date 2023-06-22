package ejercisios

import (
	"fmt"
	"strconv"
)

func RetornarDosValores(s string) (i int, ss string) {
	if num, _ := strconv.Atoi(s); num > 100 {
		i, _ := strconv.Atoi(s)
		ss := fmt.Sprintln("El valor es mayor a 100")
		return i, ss
	} else {
		i, _ := strconv.Atoi(s)
		ss := fmt.Sprintln("El valor es menor a 100")
		return i, ss
	}
}
