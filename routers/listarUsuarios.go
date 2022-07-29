package routers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/Nicolasgarcia03/Proyecto_GO/bd"
)

func ListarUsuarios(w http.ResponseWriter, r *http.Request) {
	tipoUsuario := r.URL.Query().Get("tipoUsuarios")
	page := r.URL.Query().Get("page")
	search := r.URL.Query().Get("search")

	pageTemp, err := strconv.Atoi(page)

	if err != nil {
		http.Error(w, "Debe enviar el parametro pagina como entero mayor a cero: ", http.StatusBadRequest)
		return
	}

	pag := int64(pageTemp)

	result, status := bd.LeerUsuarios(IDUsuario, pag, search, tipoUsuario)

	if status == false {
		http.Error(w, "Error al leer usuarios: ", http.StatusBadRequest)
		return
	}
	w.Header().Set("content-type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(result)
}
