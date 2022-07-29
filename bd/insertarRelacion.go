package bd

import (
	"context"
	"time"

	"github.com/Nicolasgarcia03/Proyecto_GO/models"
)

func InsertarRelacion(t models.Relacion) (bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	col := ClientMongo.Collection("relacion")
	_, err := col.InsertOne(ctx, t)

	if err != nil {
		return false, err
	}
	return true, nil
}
