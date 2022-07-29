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
	router.HandleFunc("/borroTweet", middlew.ChequeoConexion(middlew.ValidacionJWT(routers.EliminarTweet))).Methods("DELETE")

	router.HandleFunc("/subirAvatar", middlew.ChequeoConexion(middlew.ValidacionJWT(routers.SubirAvatar))).Methods("POST")
	router.HandleFunc("/obtenerAvatar", middlew.ChequeoConexion(middlew.ValidacionJWT(routers.ObtenerAvatar))).Methods("GET")
	router.HandleFunc("/subirBanner", middlew.ChequeoConexion(middlew.ValidacionJWT(routers.SubirBanner))).Methods("POST")
	router.HandleFunc("/obtenerBanner", middlew.ChequeoConexion(middlew.ValidacionJWT(routers.SubirBanner))).Methods("GET")

	router.HandleFunc("/altaRelacion", middlew.ChequeoConexion(middlew.ValidacionJWT(routers.AltaRelacion))).Methods("POST")
	router.HandleFunc("/bajaRelacion", middlew.ChequeoConexion(middlew.ValidacionJWT(routers.BajaRelacion))).Methods("DELETE")
	router.HandleFunc("/consultaRelacion", middlew.ChequeoConexion(middlew.ValidacionJWT(routers.ConsultarRelacion))).Methods("GET")

	router.HandleFunc("/listaraUsuarios", middlew.ChequeoConexion(middlew.ValidacionJWT(routers.ListarUsuarios))).Methods("GET")
	router.HandleFunc("/leerTweetsSeguidores", middlew.ChequeoConexion(middlew.ValidacionJWT(routers.LeerTweetsSeguidores))).Methods("GET")

	PORT := os.Getenv("PORT")

	if PORT == "" {
		PORT = "8080"
	}
	handler := cors.AllowAll().Handler(router)
	log.Fatal(http.ListenAndServe(":"+PORT, handler))

}
