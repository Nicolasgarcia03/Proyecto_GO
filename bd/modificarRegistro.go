package bd

import (
	"context"
	"time"

	"github.com/Nicolasgarcia03/Proyecto_GO/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func ModificarRegistro(t models.Usuario, ID string) (bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*15)
	defer cancel()

	col := ClientMongo.Collection("usuarios")

	registro := make(map[string]interface{})

	if len(t.Nombre) > 0 {
		registro["nombre"] = t.Nombre
	}
	if len(t.Apellido) > 0 {
		registro["apellido"] = t.Apellido
	}
	if len(t.Avatar) > 0 {
		registro["avatar"] = t.Avatar
	}
	if len(t.Banner) > 0 {
		registro["banner"] = t.Banner
	}
	if len(t.Biografia) > 0 {
		registro["biografia"] = t.Biografia
	}
	if len(t.Ubicacion) > 0 {
		registro["ubicacion"] = t.Ubicacion
	}
	if len(t.SitioWeb) > 0 {
		registro["sitioWeb"] = t.SitioWeb
	}
	if len(t.Email) > 0 {
		registro["email"] = t.Email
	}
	registro["fechaNacimiento"] = t.FechaNacimiento

	updateString := bson.M{
		"$set": registro,
	}

	objID, _ := primitive.ObjectIDFromHex(ID)

	filtro := bson.M{
		"_id": bson.M{
			"$eq": objID,
		},
	}
	_, err := col.UpdateOne(ctx, filtro, updateString)

	if err != nil {
		return false, err
	}
	return true, nil
}
