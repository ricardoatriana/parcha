package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

/*DevuelvoParlasSeguidores es la estructura con la que devolvemos las Parlas. Da las parlas d un usuario puede ser el mio o de cualquiera. Esto en si es como devolver esta info al http para q sea procesada por el frontend*/
type DevuelvoParlasSeguidores struct {
	ID                primitive.ObjectID `bson:"_id" json:"_id,omitempty"`
	UsarioID          string             `bson:"usuarioid" json:"userId,omitempty"` //Id lo pusimos en mayus por convencion
	UsuarioRelacionID string             `bson:"usuariorelacionid" json:"userRelationID,omitempty"`
	Parla             struct {
		Mensaje string    `bson:"mensaje" json:"mensaje,omitempty"` //este es el mismo tweet
		Fecha   time.Time `bson:"fecha" json:"fecha,omitempty"`
		ID      string    `bson:"_id" json:"_id,omitempty"` //ID del tweet
	}
}
