package ejerinterfaces

import (
	"fmt"

	"github.com/FernandoMendoza12/godesde0/interfaces"
)

func HumanosRespirando(hu interfaces.Humano) {
	hu.Respirar()
	fmt.Printf("Soy un o una %s, y estoy respirando \n", hu.Sexo())
}
