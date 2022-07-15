package handlers

import (
	"log"
	"net/http"
	"os"

	"github.com/Nicolasgarcia03/Proyecto_GO/middlew"
	"github.com/Nicolasgarcia03/Proyecto_GO/routers"
	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

/*Manejadore para el puerto*/
func Manejadores() {
	router := mux.NewRouter()

	router.HandleFunc("/registro", middlew.ChequeoConexion(routers.Registro)).Methods("POST")

	PORT := os.Getenv("PORT")

	if PORT == "" {
		PORT = "8080"
	}
	handler := cors.AllowAll().Handler(router)
	log.Fatal(http.ListenAndServe(":"+PORT, handler))

}
