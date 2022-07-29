package models

type Relacion struct {
	UsuarioID         string `bson:"usuarioID" json:"usuarioID"`
	UsuarioRelacionID string `bson:"usuarioRelacionID" json:"usuarioRelacionID"`
}
