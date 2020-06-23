package bd

import (
	"context" //en el paquete bd todas usan context, casi todas usan time
	"time"

	"github.com/ricardoatriana/parcha/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

/*InsertoParla graba parla en BD*/
func InsertoParla(t models.GraboParla) (string, bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second) // como vamos hacer operaciones de grabacion, lectura de la bd, queremos q nose quede colgada la bd asi q creamos el contexto cxt
	defer cancel()                                                           //es una instruccion q se setea a comienzo pero se ejectua omo ultima instruccion de la funcion. el cancel cancela el widthtimeout, si no se hace asi, dnetro de la gran caja de contexto se van creando cajitas de operiaciones y si no se dan de baja van ocuapdno mucho espacio en el contexto

	db := MongoCN.Database("parcha")
	col := db.Collection("parla")

	registro := bson.M{
		"userid":  t.UserID, //grabamos el id del usuario q deja el msj
		"mensaje": t.Mensaje,
		"fecha":   t.Fecha,
	}

	result, err := col.InsertOne(ctx, registro) //result y error son los 2 parametros q traen la coleccion, q tiene mi contexto y va insertar el registro
	if err != nil {
		return "", false, err //tenemos 3 campos a regresar en uno va el ID, en otro el bool y el otro va el error. Como aqui es un error vamos a devolver un string vacio que se puede hacer asi return "", o por convencion string("")
	}

	objID, _ := result.InsertedID.(primitive.ObjectID) //La func de primitive devuelve un objID y un parametro q no nos interesa. Del json/bson q devuelve el InsertOne extrae la clave del ultimo campo insertado y ahi obtiene el objectID.
	return objID.String(), true, nil                   //se convierte el ObjID en string. Estos se puede hacer con .String o .Hex

}
