package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

/*Usuario es el modelo de usuario de la base de MOngoDB */
type Usuario struct {
	ID              primitive.ObjectID `bson:"_id" json:"id,omitempty"`        //el ID de mongo no es un num, no es un long esun objeto binario llamado ObjectID, es un slide de bytes. Recordar q en BD todo se graba en formato bson. Los ID de Mongo se graban _id
	Nombre          string             `bson:"nombre" json:"nombre,omitempty"` //uno son los datos de entrada a la base y los otros son los datos de salida al navegador
	Apellido        string             `bson:"apellido" json:"apellido,omitempty"`
	FechaNacimiento time.Time          `bson:"fechaNacimiento" json:"fechaNacimiento,omitempty"`
	Email           string             `bson:"email" json:"email,omitempty"` //omitempty hace qe no devuela por el navegado
	Clave           string             `bson:"clave" json:"clave,omitempty"`
	Avatar          string             `bson:"avatar" json:"avatar,omitempty"`
	Banner          string             `bson:"banner" json:"banner,omitempty"`
	Biografia       string             `bson:"biografia" json:"biografia,omitempty"`
	Ubicacion       string             `bson:"ubicacion" json:"ubicacion,omitempty"`
	SitioWeb        string             `bson:"sitioWeb" json:"sitioWeb,omitempty"`
}
