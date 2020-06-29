package bd

import (
	"context"
	"fmt"
	"time"

	"github.com/ricardoatriana/parcha/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

/*LeoUsuariosTodos lee los usuarios registrados en el sistema, si recibe R en quienes trae solo los q relaciona conmigo */
func LeoUsuariosTodos(ID string, page int64, search string, tipo string) ([]*models.Usuario, bool) { //ID es el q esta leyedo los usuarios, search nos va a permitir filtrar por un termino, tipo es si buscamos dentro de nuestra red o dentro de toda la app. Devuelve un slice porque ahora no se va adevolver un modelo de usuario, va a devolver n modelos de usuarios
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*15)
	defer cancel()

	db := MongoCN.Database("parcha")
	col := db.Collection("usuarios")

	var results []*models.Usuario

	findOptions := options.Find()
	findOptions.SetSkip((page - 1) * 20)
	findOptions.SetLimit(20)

	query := bson.M{ //le estamos dando una condicion a la tabla usuarios, q el campo nombre sea de tipo bson
		"nombre": bson.M{"$regex": `(?i)` + search}, //Va a buscar dentro de los strings no importa si son mayusculas o minusculas
	}

	//ejecutamos el find
	cur, err := col.Find(ctx, query, findOptions) //cuando no es un findone el resultado me lo de en un cursor
	if err != nil {
		fmt.Println(err.Error())
		return results, false
	}

	// este lo q hace es por ej: Aparece primero Pedro, hace una iteracion, lego aparece Maria, hace de nuevo otra iteracion aparece mas
	var encontrado, incluir bool //luego vamos a hacer un for, con este vamos a recorrer cursor
	for cur.Next(ctx) {          //El next permite avanzar al siguiente registro le pasa el contexto de parametro
		var s models.Usuario
		err := cur.Decode(&s) //lo guardamos en la posicion de memoria donde he creado un modelo usuario
		if err != nil {
			fmt.Println(err.Error())
			return results, false
		}

		var r models.Relacion //consulta la relacion de este usuario co los demas
		r.UsuarioID = ID
		r.UsuarioRelacionID = s.ID.Hex() // de los usuarios extraigo el ID el ObjID

		incluir = false //por cada usuario q esta apareciendo en la iteracion yo voy a querer saber si lo quiero incluir en ua respuesta o no

		encontrado, err = ConsultoRelacion(r)
		if tipo == "new" && encontrado == false { //encontrado false quiere decir es un usuario al cual yo no sigo
			incluir = true
		}
		if tipo == "follow" && encontrado == true { //encontrado false quiere decir es un usuario al cual yo no sigo
			incluir = true
		}
		if r.UsuarioRelacionID == ID {
			incluir = false
		}

		if incluir == true { //cuadno se decodifca el cursor se trae todos los elementos, asi q los seteamos
			s.Clave = ""
			s.Biografia = ""
			s.SitioWeb = ""
			s.Ubicacion = ""
			s.Banner = ""
			s.Email = ""

			results = append(results, &s) //este append adiciona los datos al slice q vamos a devolver

		}
	}
	err = cur.Err()
	if err != nil {
		fmt.Println(err.Error())
		return results, false
	}
	cur.Close(ctx)
	return results, true
}
