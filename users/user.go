package users

import (
	"fmt"
	"github.com/386291john/godesde0/modelos"
	"time"
)

func AltaUsua() {
	u := new(modelos.User)
	u.AddUser(10, "Pablo", time.Now(), true)
	fmt.Println(u)
}
