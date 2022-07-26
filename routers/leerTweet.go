package routers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/Nicolasgarcia03/Proyecto_GO/bd"
)

func LeerTweet(w http.ResponseWriter, r *http.Request) {
	ID := r.URL.Query().Get("id")
	if (len(ID) < 1) && (len(r.URL.Query().Get("pagina")) < 1) {
		http.Error(w, "el id y la pagina son obligatorios ", http.StatusBadRequest)
		return
	}
	pagina, err := strconv.Atoi(r.URL.Query().Get("pagina"))

	if err != nil {
		http.Error(w, "la pagina debe ser mayor a 0 ", http.StatusBadRequest)
		return
	}

	pag := int64(pagina)

	respuesta, correcto := bd.LeerTweet(ID, pag)

	if correcto == false {
		http.Error(w, "Hubo un error leyendo twweets ", http.StatusBadRequest)
		return
	}

	w.Header().Set("content-type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(respuesta)
}
