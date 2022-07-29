package bd

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func BorroTweet(ID string, UsuarioID string) error {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*15)
	defer cancel()

	col := ClientMongo.Collection("tweet")

	objID, _ := primitive.ObjectIDFromHex(ID)

	condicion := bson.M{
		"_id":       objID,
		"usuarioid": UsuarioID,
	}

	_, err := col.DeleteOne(ctx, condicion)

	return err
}
