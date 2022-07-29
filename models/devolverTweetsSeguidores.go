package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type DevolverTweetsSeguidores struct {
	ID                primitive.ObjectID `bson:"_id" json:"_id,omitempty"`
	UsuarioID         string             `bson:"usuarioID" json:"usuarioID,omitempty"`
	UsuarioRelacionID string             `bson:"usuarioRelacionID" json:"usuarioRelacionID,omitempty"`
	Tweet             struct {
		Mensaje string    `bson:"mensaje" json:"mensaje,omitempty"`
		Fecha   time.Time `bson:"fecha" json:"fecha,omitempty"`
		ID      string    `bson:"_id" json:"_id,omitempty"`
	}
}
