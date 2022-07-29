package bd

import (
	"context"
	"fmt"
	"time"

	"github.com/Nicolasgarcia03/Proyecto_GO/models"
	"go.mongodb.org/mongo-driver/bson"
)

func ConsultarRelacion(t models.Relacion) (bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	col := ClientMongo.Collection("relacion")

	condicion := bson.M{
		"usuarioid":         t.UsuarioID,
		"usuariorelacionid": t.UsuarioRelacionID,
	}

	var resultado models.Relacion
	fmt.Println(resultado)
	err := col.FindOne(ctx, condicion).Decode(&resultado)

	if err != nil {
		fmt.Println(err.Error())
		return false, err
	}
	return true, nil
}
