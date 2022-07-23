package bd

import (
	"github.com/Nicolasgarcia03/Proyecto_GO/models"
	"golang.org/x/crypto/bcrypt"
)

/*Login*/
func Login(email string, password string) (models.Usuario, bool) {
	usu, existe, _ := ExisteUsuario(email)

	if existe == false {
		return usu, false
	}

	passByte := []byte(password)
	passBD := []byte(usu.Password)
	err := bcrypt.CompareHashAndPassword(passBD, passByte)
	if err != nil {
		return usu, false
	}
	return usu, true
}
