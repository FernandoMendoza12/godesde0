package goroutines

import (
	"fmt"
	"strings"
	"time"
)

func MiNombreLento(name string, ch1 chan bool) {
	letras := strings.Split(name, "")
	for _, v := range letras {
		time.Sleep(1000 * time.Millisecond)
		fmt.Println(v)
	}
	ch1 <- true
}
