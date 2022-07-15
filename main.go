package main

import (
	"github.com/Nicolasgarcia03/Proyecto_GO/bd"
	"github.com/Nicolasgarcia03/Proyecto_GO/handlers"
)

func main() {
	conexionDBMongo()
	manejadores()
}

func conexionDBMongo() {
	bd.ConectarBD()
}

func manejadores() {
	handlers.Manejadores()
}
