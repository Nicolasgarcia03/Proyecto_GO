package routers

import (
	"encoding/json"
	"net/http"

	"github.com/Nicolasgarcia03/Proyecto_GO/bd"
	"github.com/Nicolasgarcia03/Proyecto_GO/models"
)

/*Registro*/
func Registro(w http.ResponseWriter, r *http.Request) {
	var t models.Usuario
	err := json.NewDecoder(r.Body).Decode(&t)

	if err != nil {
		http.Error(w, "Error en los datos recibidos"+err.Error(), 400)
		return
	}

	if len(t.Email) == 0 {
		http.Error(w, "El Email de Usuario es requerido", 400)
		return
	}

	if len(t.Password) < 6 {
		http.Error(w, "La password debe contener mas de 6 caracteres", 400)
		return
	}

	_, encontrado, _ := bd.ExisteUsuario(t.Email)

	if encontrado == true {
		http.Error(w, "Ya existe un usuario con ese email", 400)
		return
	}

	_, status, err := bd.InsertarRegistros(t)

	if err != nil {
		http.Error(w, "Ocurrio un error al insertar el registro"+err.Error(), 400)
		return
	}

	if status == false {
		http.Error(w, "Fallo la insercion del registr de usuario", 400)
		return
	}

	w.WriteHeader(http.StatusCreated)
}
