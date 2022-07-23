package routers

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/Nicolasgarcia03/Proyecto_GO/bd"
	"github.com/Nicolasgarcia03/Proyecto_GO/jwt"
	"github.com/Nicolasgarcia03/Proyecto_GO/models"
)

/*login*/
func Login(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("content-type", "application/json")
	var t models.Usuario

	err := json.NewDecoder(r.Body).Decode(&t)

	if err != nil {
		http.Error(w, "Usuarios y/o Contraseña invalido"+err.Error(), 400)
		return
	}
	if len(t.Email) == 0 {
		http.Error(w, "Email de usuario requerido", 400)
	}

	doc, existe := bd.Login(t.Email, t.Password)
	if existe == false {
		http.Error(w, "Usuario y/o Contaseña invalido", 400)
	}

	jwtkey, err := jwt.GeneroJWT(doc)

	if err != nil {
		http.Error(w, "Ocurrio un Error intentando generar el Token"+err.Error(), 400)
		return
	}

	resp := models.RespuestaLogin{
		Token: jwtkey,
	}

	w.Header().Set("content-type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(resp)

	expiracionTime := time.Now().Add(24 * time.Hour)
	http.SetCookie(w, &http.Cookie{
		Name:    "Token",
		Value:   jwtkey,
		Expires: expiracionTime,
	})
}
