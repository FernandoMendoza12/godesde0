package deferPanic

import (
	"fmt"
)

func VemosDefer() {
	fmt.Println("Este es el primer mensaje")
	defer fmt.Println("Este es el mensaje final")

	fmt.Println("Este es el segundo mensaje")
}
