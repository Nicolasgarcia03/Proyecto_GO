package middlew

import (
	"net/http"

	"github.com/Nicolasgarcia03/Proyecto_GO/bd"
)

func ChequeoConexion(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if bd.ClientMongo == nil {
			http.Error(w, "Conexion perdida", 500)
			return
		}
		next.ServeHTTP(w, r)
	}
}
