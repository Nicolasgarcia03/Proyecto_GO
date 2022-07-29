package bd

import (
	"context"
	"time"

	"github.com/Nicolasgarcia03/Proyecto_GO/models"
	"go.mongodb.org/mongo-driver/bson"
)

func LeerTweetsSeguidores(ID string, pagina int) ([]models.DevolverTweetsSeguidores, bool) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	col := ClientMongo.Collection("relacion")

	skip := (pagina - 1) * 20

	condiciones := make([]bson.M, 0)
	condiciones = append(condiciones, bson.M{"$match": bson.M{"usuarioID": ID}})
	condiciones = append(condiciones, bson.M{
		"$lookup": bson.M{
			"from":         "tweet",
			"localField":   "usuarioRelacionID",
			"foreignField": "UsuarioID",
			"as":           "tweet",
		},
	})
	condiciones = append(condiciones, bson.M{"$unwind": "$tweet"})
	condiciones = append(condiciones, bson.M{"$sort": bson.M{"fecha": -1}})
	condiciones = append(condiciones, bson.M{"$skip": skip})
	condiciones = append(condiciones, bson.M{"$limit": 20})

	cursor, err := col.Aggregate(ctx, condiciones)
	var resultado []models.DevolverTweetsSeguidores
	err = cursor.All(ctx, &resultado)

	if err != nil {
		return resultado, false
	}

	return resultado, true
}
