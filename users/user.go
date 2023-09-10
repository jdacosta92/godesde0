package users

import (
	"fmt"
	"time"

	"github.com/jdacosta92/godesde0/modelos"
)

func AltaUsuario() {
	u := new(modelos.User)
	u.AddUser(1, "Julian", time.Now(), true)
	fmt.Println(u)
	fmt.Println()
	fmt.Println(u.Id, " ", u.Name, " ", u.CreatedAt, " ", u.Status)
}
