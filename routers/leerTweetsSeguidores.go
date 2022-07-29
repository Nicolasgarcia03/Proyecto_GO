package routers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/Nicolasgarcia03/Proyecto_GO/bd"
)

func LeerTweetsSeguidores(w http.ResponseWriter, r *http.Request) {

	if len(r.URL.Query().Get("pagina")) < 1 {
		http.Error(w, "el parametro pagina es obligatorio", http.StatusBadRequest)
		return
	}
	pagina, err := strconv.Atoi(r.URL.Query().Get("pagina"))
	if err != nil {
		http.Error(w, "el parametro pagina debe ser enviado como entero mayor a 0 ", http.StatusBadRequest)
		return
	}

	respuesta, status := bd.LeerTweetsSeguidores(IDUsuario, pagina)

	if status == false {
		http.Error(w, "error al leer tweets ", http.StatusBadRequest)
		return
	}
	w.Header().Set("content-type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(respuesta)
}
