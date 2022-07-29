package bd

import (
	"context"
	"fmt"
	"time"

	"github.com/Nicolasgarcia03/Proyecto_GO/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func LeerUsuarios(ID string, page int64, search string, tipo string) ([]*models.Usuario, bool) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	col := ClientMongo.Collection("usuarios")

	var resultados []*models.Usuario

	findOptions := options.Find()
	findOptions.SetSkip((page - 1) * 20)
	findOptions.SetLimit(20)

	query := bson.M{
		"nombre": bson.M{
			"$regex": `(?i)` + search,
		},
	}

	cursor, err := col.Find(ctx, query, findOptions)

	if err != nil {
		fmt.Println(err.Error())
		return resultados, false
	}
	var encontrado, incluir bool
	for cursor.Next(ctx) {
		var s models.Usuario
		err := cursor.Decode(&s)
		if err != nil {
			fmt.Println(err.Error())
			return resultados, false
		}

		var r models.Relacion
		r.UsuarioID = ID
		r.UsuarioRelacionID = s.ID.Hex()

		incluir = false

		encontrado, err = ConsultarRelacion(r)

		if tipo == "new" && encontrado == false {
			incluir = true
		}

		if tipo == "follow" && encontrado == true {
			incluir = true
		}

		if r.UsuarioRelacionID == ID {
			incluir = false
		}

		if incluir == true {
			s.Password = ""
			s.Biografia = ""
			s.SitioWeb = ""
			s.Ubicacion = ""
			s.Banner = ""
			s.Email = ""

			resultados = append(resultados, &s)

		}
	}

	err = cursor.Err()
	if err != nil {
		fmt.Println(err.Error())
		return resultados, false
	}

	cursor.Close(ctx)

	return resultados, true
}
