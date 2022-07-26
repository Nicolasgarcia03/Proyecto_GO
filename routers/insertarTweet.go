package routers

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/Nicolasgarcia03/Proyecto_GO/bd"
	"github.com/Nicolasgarcia03/Proyecto_GO/models"
)

func InsertarTweet(w http.ResponseWriter, r *http.Request) {
	var mensaje models.Tweet

	err := json.NewDecoder(r.Body).Decode(&mensaje)

	registro := models.GraboTweet{
		UsuarioID: IDUsuario,
		Mensaje:   mensaje.Mensaje,
		Fecha:     time.Now(),
	}
	_, status, err := bd.InsertarTweet(registro)

	if err != nil {
		http.Error(w, "Ocurrio un error al intentar insertar el tweet, reintente nuevamente: "+err.Error(), http.StatusBadRequest)
		return
	}

	if status == false {
		http.Error(w, "No se logro insertar el tweet, reintente nuevamente: ", http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusCreated)
}
