package bd

import (
	"context" //en el paquete bd todas usan context, casi todas usan time
	"time"

	"github.com/ricardoatriana/parcha/models"
	"go.mongodb.org/mongo-driver/bson"
)

/*ChequeoYaExisteUsuario recibe un email de parametro y chequea si ya esta en la BD*/
func ChequeoYaExisteUsuario(email string) (models.Usuario, bool, string) { //el models es usuario porque una vez encontro usuario en la bd me devuelve todo el registro
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	db := MongoCN.Database("parcha")
	col := db.Collection("usuarios")

	condicion := bson.M{"email": email} //condiciones en Mongo, el main tiene ser igual a lo q se recibe de parametro, con una var tipo bson

	var resultado models.Usuario //modelando el usuario en una variable resultado

	err := col.FindOne(ctx, condicion).Decode(&resultado)
	ID := resultado.ID.Hex() //Se esta convirtiendo el resultado a string hexadecimal
	if err != nil {
		return resultado, false, ID
	}
	return resultado, true, ID
}
