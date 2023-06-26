package structs

import (
	"fmt"
	"time"

	"github.com/FernandoMendoza12/godesde0/modelos"
)

func AltaUsuario() {
	u := new(modelos.User)
	u.AddUser(1, "Fernando", time.Now(), true)
	fmt.Printf("The struct have ID: %d and the name is : %s, createdAt : %s and is active: %v", u.Id, u.Name, u.CreatedAt, u.Status)
}
