package bd

import (
	"context"
	"time"

	"github.com/ricardoatriana/parcha/models"
	"go.mongodb.org/mongo-driver/bson"
)

/*LeoParlasSeguidores lee las parlas de mis seguidores*/
func LeoParlasSeguidores(ID string, pagina int) ([]models.DevuelvoParlasSeguidores, bool) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*15)
	defer cancel()

	db := MongoCN.Database("parcha")
	col := db.Collection("relacion")
	//De la tabla relacion se va a sacar la info basica de con quien estoy relacionado y luego vamos a unir esta tabla con las parlas
	skip := (pagina - 1) * 20 //vamos solo a traer 20 tweets

	condiciones := make([]bson.M, 0)
	condiciones = append(condiciones, bson.M{"$match": bson.M{"usuarioid": ID}}) //que haga un match y la condiciones es que usuarioid sa igual al ID q vino por parametro
	condiciones = append(condiciones, bson.M{                                    // la instruccion lookup permite unir 2 tablas, una vez q encontre y filtre mi tabla relacion por mi userid, va a atraer todos esos registros, y voy hacer q ese resultado este unido a la tabla parlas
		"$lookup": bson.M{
			"from":         "parla",             //voy a unir relacion con mi tabla parla
			"localField":   "usuariorelacionid", // este es el campopor el cual vamos a unir. usuarioid es el nuestro ahora vamos aunir la tabla parla en base a nuestros usuarios con los q estamos relacionados porque tenemos q traer las parlas de ellos
			"foreignField": "userid",
			"as":           "parla", //un alias de como queremos llamar esa nueva tabla, pues le llamamos parla, igual q la madre
		}})
	condiciones = append(condiciones, bson.M{"$unwind": "$parla"})                //por ejemplo si estoy unido a pepe,y no usaramos unwind apareceria un doc inicial con los datos de pepe y por debajo un array de docs con todos las parlas de el. El unwind permite es q todos los docs vengan exactamente iguales, entonces los datos de pepe se van a repetir en todos las parlas y va a aestar su tweet en el mismo doc
	condiciones = append(condiciones, bson.M{"$sort": bson.M{"parla.fecha": -1}}) //El ordenado de menor a mayor es 1, pero si es de ultimo a primer el mas antiguo al final es -1
	condiciones = append(condiciones, bson.M{"$skip": skip})
	condiciones = append(condiciones, bson.M{"$limit": 20})

	cursor, err := col.Aggregate(ctx, condiciones) //esto se ejecuta en la bd con todas las condiciones y crea un curus
	var result []models.DevuelvoParlasSeguidores
	err = cursor.All(ctx, &result) //decodifica todo en result
	if err != nil {
		return result, false
	}
	return result, true
}
