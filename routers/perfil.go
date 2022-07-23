package routers

import (
	"encoding/json"
	"net/http"

	"github.com/Nicolasgarcia03/Proyecto_GO/bd"
)

func ObtenerPerfil(w http.ResponseWriter, r *http.Request) {
	ID := r.URL.Query().Get("id")
	if len(ID) < 1 {
		http.Error(w, "Debe enviar el parametro ID", http.StatusBadRequest)
		return
	}
	perfil, err := bd.ObtenerPerfil(ID)

	if err != nil {
		http.Error(w, "Ocurrio un Error buscando el perfil"+err.Error(), http.StatusBadRequest)
		return
	}
	w.Header().Set("content-type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(perfil)
}
