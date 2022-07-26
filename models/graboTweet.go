package models

import "time"

type GraboTweet struct {
	UsuarioID string    `bson:"usuarioID" json:"usuarioID,omitempty"`
	Mensaje   string    `bson:"mensaje" json:"mensaje,omitempty"`
	Fecha     time.Time `bson:"fecha" json:"fecha,omitempty"`
}
