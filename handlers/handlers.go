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
	router.HandleFunc("/login", middlew.ChequeoConexion(routers.Login)).Methods("POST")
	router.HandleFunc("/verperfil", middlew.ChequeoConexion(middlew.ValidacionJWT(routers.ObtenerPerfil))).Methods("GET")
	router.HandleFunc("/modificarregistro", middlew.ChequeoConexion(middlew.ValidacionJWT(routers.ModificarPerfil))).Methods("PUT")
	router.HandleFunc("/tweet", middlew.ChequeoConexion(middlew.ValidacionJWT(routers.InsertarTweet))).Methods("POST")
	router.HandleFunc("/leerTweet", middlew.ChequeoConexion(middlew.ValidacionJWT(routers.LeerTweet))).Methods("GET")

	PORT := os.Getenv("PORT")

	if PORT == "" {
		PORT = "8080"
	}
	handler := cors.AllowAll().Handler(router)
	log.Fatal(http.ListenAndServe(":"+PORT, handler))

}
