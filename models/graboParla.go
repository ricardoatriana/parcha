package models

import "time"

/*GraboParla es el fomrato o la estructura que tendra nuestro Parla n la BD*/
type GraboParla struct {
	UserID  string    `bson:"userid" json:"userid,omitempty"` //en la base de datos se va a llamar como se llama en bson y tiene un representacion json
	Mensaje string    `bson:"mensaje" json:"mensaje,omitempty"`
	Fecha   time.Time `bson:"fecha" json:"fecha,omitempty"`
}
