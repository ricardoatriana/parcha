package bd

import (
	"context"
	"fmt"
	"time"

	"github.com/ricardoatriana/parcha/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

/*BuscoPerfil busca perfil en la BD*/
func BuscoPerfil(ID string) (models.Usuario, error) { //Va a venir como parametro a modo de peticion GET
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*15)
	defer cancel()

	db := MongoCN.Database("parcha")
	col := db.Collection("usuarios")

	var perfil models.Usuario                 // no puedo ir a bsucar un usuario a Monog diciendole q el string tiene q coincidir con el ObjectID, porque el ID en la coleccion de usarios es de OBjectID, son objetos diferentes
	objID, _ := primitive.ObjectIDFromHex(ID) // aqui se realiza de cambio

	condicion := bson.M{ //la condicion de busqueda es un bsonM y asi bson M, tiene q coincidir los 2 objetos ids. Haciendo q se busq en un objeto igual a de la base de mongo
		"_id": objID,
	}
	err := col.FindOne(ctx, condicion).Decode(&perfil) //de lo q encuentra con el find haga un decode
	perfil.Clave = ""                                  //asi no vuelve a mi como valor, y no nos enteramos de la clave
	if err != nil {                                    //si el error q dio el findone es distinto de nil
		fmt.Println("Registro no encontrado" + err.Error())
		return perfil, err
	}
	return perfil, nil //si encontro un perfil del usuario hacemos un return de nil, no deuvelve err
}
