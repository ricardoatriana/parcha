package bd

import (
	"context"
	"time"

	"github.com/ricardoatriana/parcha/models"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

/*InsertoRegistro es la parada final con la BD para insertar los datos del usuario */
func InsertoRegistro(u models.Usuario) (string, bool, error) { //en los parametros va primero el nombre q le queremos dar y luego el tipo de dato q es usario de tipo json
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second) // como vamos hacer operaciones de grabacion, lectura de la bd, queremos q nose quede colgada la bd asi q creamos el contexto cxt
	defer cancel()                                                           //es una instruccion q se setea a comienzo pero se ejectua omo ultima instruccion de la funcion. el cancel cancela el widthtimeout, si no se hace asi, dnetro de la gran caja de contexto se van creando cajitas de operiaciones y si no se dan de baja van ocuapdno mucho espacio en el contexto

	db := MongoCN.Database("parcha")
	col := db.Collection("usuarios")     // hay 3 colecciones; usuario, twitter y relacion. Apunta a la coleccion de usuarios de la bd de parcha
	u.Clave, _ = EncriptarClave(u.Clave) //no se pueden grabar claves literales en las bd's es un tema de seguridad, por es se encripta

	result, err := col.InsertOne(ctx, u) // el insertone es para insertar un solo registro, el ctx es para q no dure mas de 15s. El result es un doc JSON
	if err != nil {
		return "", false, err
	}
	ObjID, _ := result.InsertedID.(primitive.ObjectID) // Esta es la forma de cambiar el objectID a string luego de haber usado el insertedOne
	return ObjID.String(), true, nil
}
