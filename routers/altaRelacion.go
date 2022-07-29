package routers

import (
	"net/http"

	"github.com/Nicolasgarcia03/Proyecto_GO/bd"
	"github.com/Nicolasgarcia03/Proyecto_GO/models"
)

func AltaRelacion(w http.ResponseWriter, r *http.Request) {
	ID := r.URL.Query().Get("id")
	if len(ID) < 1 {
		http.Error(w, "El parametro ID es oblogatorio", http.StatusBadRequest)
		return
	}

	var t models.Relacion

	t.UsuarioID = IDUsuario
	t.UsuarioRelacionID = ID

	status, err := bd.InsertarRelacion(t)
	if err != nil {
		http.Error(w, "Ha ocurrido un Error: "+err.Error(), http.StatusBadRequest)
		return
	}

	if status == false {
		http.Error(w, "No se ha logrado insertar la relacion ", http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusCreated)
}
