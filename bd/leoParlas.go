package bd

import (
	"context" //en el paquete bd todas usan context, casi todas usan time
	"time"

	"github.com/ricardoatriana/parcha/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options" //define opciones para poder filtrar y darle un comportamiento a la BD
)

/*LeoParlas lee las parlas de un perfil*/
func LeoParlas(ID string, pagina int64) ([]*models.DevuelvoParlas, bool) { //ID del usuario del cual voy a leer las parlas. [] las parlas son un vector q devuelve varias parlas de un solo golpe puede ser 15 o 20 parlas
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*15)
	defer cancel()

	db := MongoCN.Database("parcha")
	col := db.Collection("parla")

	var resultados []*models.DevuelvoParlas //aqui es donde vamos a grabar los resultados, y ahora creamos la condicion por la cual buscamos en la BD
	condicion := bson.M{                    //voy a leer mi tabla parla donde el userid sea el q me llego por parametro
		"userid": ID, // es el ID q le pasamos desde el paquete router a esta rutina de BD
	}

	opciones := options.Find()
	opciones.SetLimit(20)                               //trae 20 parlas
	opciones.SetSort(bson.D{{Key: "fecha", Value: -1}}) //viene ordenado por fecha, de forma -1 q significa descendente
	opciones.SetSkip((pagina - 1) * 20)                 // entonces cuadno se quiera traer las parlas de la pag1, sera asi. 1-1=0*20, entonces no salta ningunn doc. Cuando envie la pag2, sera 2-1=1*20, salta los primeros 20 registros

	cursor, err := col.Find(ctx, condicion, opciones) // cursor es como si fuera una tabla donde se guardan los resultado
	if err != nil {
		return resultados, false
	}

	for cursor.Next(context.TODO()) { //por cada iteraccion crea una var llamada registro, cuando vuevle a iterar vuelve a cargar registro en memoria, hace un decode de cursor en el registro, si todo sale bien a resultados le agrega con el append el mismo registro
		var registro models.DevuelvoParlas //para trabajar con cada parla en particular
		err := cursor.Decode(&registro)
		if err != nil {
			return resultados, false
		}
		resultados = append(resultados, &registro) //
	}
	return resultados, true
}
