package routers

import (
	"encoding/json"
	"net/http"

	"github.com/Nicolasgarcia03/Proyecto_GO/bd"
	"github.com/Nicolasgarcia03/Proyecto_GO/models"
)

func ConsultarRelacion(w http.ResponseWriter, r *http.Request) {
	ID := r.URL.Query().Get("id")
	var t models.Relacion
	t.UsuarioID = IDUsuario
	t.UsuarioRelacionID = ID

	var respuesta models.RespuestaConsultarRelacion

	status, err := bd.ConsultarRelacion(t)

	if err != nil || status == false {
		respuesta.Status = false
	} else {
		respuesta.Status = true
	}

	w.Header().Set("content-type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(respuesta)
}
