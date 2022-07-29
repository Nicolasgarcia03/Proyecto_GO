package bd

import (
	"context"
	"time"

	"github.com/Nicolasgarcia03/Proyecto_GO/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func InsertarTweet(t models.GraboTweet) (string, bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*15)

	defer cancel()

	col := ClientMongo.Collection("tweet")

	registro := bson.M{
		"usuarioid": t.UsuarioID,
		"mensaje":   t.Mensaje,
		"fecha":     t.Fecha,
	}

	resultado, err := col.InsertOne(ctx, registro)

	if err != nil {
		return string(""), false, err
	}

	objID, _ := resultado.InsertedID.(primitive.ObjectID)

	return objID.String(), true, nil
}
