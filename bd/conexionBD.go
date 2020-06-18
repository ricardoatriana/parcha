package bd

import (
	"context"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

/*MongoCN is el objeto de conexion a la BD*/
var MongoCN = ConectarBD() //var a exportar por eso comienza en mayusucla, va a ser exportada a todos los rchivos q usan bd, si es de uso interno es en minuscula
var clientOptions = options.Client().ApplyURI("mongodb+srv://ricardoatriana:ricardo2@cluster0-igdae.mongodb.net/parcha?retryWrites=true&w=majority")

/*ConectarBD es la funcion q me permite conectar la base de datos*/
func ConectarBD() *mongo.Client {
	client, err := mongo.Connect(context.TODO(), clientOptions) //contexto no sirve para comunicar info entre ejecuccion y ejecucion, y y admeas nos permite setear como por ejemplo un timeOut de 15s, para eviat
	if err != nil {
		log.Fatal(err.Error()) //Error() agarra el error y lo convierte a string
		return client
	}
	err = client.Ping(context.TODO(), nil) //Cuado la var se usa por 1ra vez como linea15 se usa :, aqui como es la 2da vez no hace falta
	if err != nil {
		log.Fatal(err.Error())
		return client
	}
	log.Println("Conexion existosa con la BD")
	return client
}

/*ChequeoConnection es el Ping a la BD*/
func ChequeoConnection() int {
	err := MongoCN.Ping(context.TODO(), nil)
	if err != nil {
		return 0
	}
	return 1
}
