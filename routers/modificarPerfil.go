package routers

import (
	"encoding/json"
	"net/http"

	"github.com/Nicolasgarcia03/Proyecto_GO/bd"
	"github.com/Nicolasgarcia03/Proyecto_GO/models"
)

func ModificarPerfil(w http.ResponseWriter, r *http.Request) {
	var t models.Usuario
	err := json.NewDecoder(r.Body).Decode(&t)

	if err != nil {
		http.Error(w, "Datos incorrectos "+err.Error(), http.StatusBadRequest)
		return
	}

	var status bool
	status, err = bd.ModificarRegistro(t, IDUsuario)

	if err != nil {
		http.Error(w, "ocurrio un error al intentar modificar el registro "+err.Error(), http.StatusBadRequest)
		return
	}

	if status == false {
		http.Error(w, "No se logro modificar el registro "+err.Error(), http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusCreated)
}
