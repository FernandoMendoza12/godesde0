package funciones

import "fmt"

func Exponencia(i int) {
	if i > 1000000 {
		return
	}
	fmt.Println(i)
	Exponencia(i * 2)
}
