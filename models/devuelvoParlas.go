package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

/*DevuelvoParlas es la estructura con la que devolvemos las Parlas. Da las parlas d un usuario puede ser el mio o de cualquiera. Esto en si es como devolver esta info al http para q sea procesada por el frontend*/
type DevuelvoParlas struct {
	ID      primitive.ObjectID `bson:"_id" json:"_id,omitempty"`
	UserID  string             `bson:"userid" json:"userId,omitempty"` //Id lo pusimos en mayus por convencion
	Mensaje string             `bson:"mensaje" json:"mensaje,omitempty"`
	Fecha   time.Time          `bson:"fecha" json:"fecha,omitempty"`
}
