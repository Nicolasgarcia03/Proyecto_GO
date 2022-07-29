package routers

import (
	"io"
	"net/http"
	"os"

	"github.com/Nicolasgarcia03/Proyecto_GO/bd"
)

func ObtenerBanner(w http.ResponseWriter, r *http.Request) {
	ID := r.URL.Query().Get("id")

	if len(ID) < 1 {
		http.Error(w, "El Id es obligatorio", http.StatusBadRequest)
		return
	}
	perfil, err := bd.ObtenerPerfil(ID)

	if err != nil {
		http.Error(w, "Usuario no encontrado ", http.StatusBadRequest)
		return
	}

	openFile, err := os.Open("uploads/avatars/" + perfil.Avatar)
	if err != nil {
		http.Error(w, "Imagen no encontrada ", http.StatusBadRequest)
		return
	}

	_, err = io.Copy(w, openFile)

	if err != nil {
		http.Error(w, "Error al copiar la imagen ", http.StatusBadRequest)
	}
}
