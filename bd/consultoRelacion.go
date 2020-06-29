package bd

import (
	"context"
	"fmt"
	"time"

	"github.com/ricardoatriana/parcha/models"
	"go.mongodb.org/mongo-driver/bson"
)

/*ConsultoRelacion consulta la relacion entre 2 usuarios. Lo q nos interesa saber es si hay relacion*/
func ConsultoRelacion(t models.Relacion) (bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*15)
	defer cancel()

	db := MongoCN.Database("parcha")
	col := db.Collection("relacion")

	condicion := bson.M{ //condicion de busqueda
		"usuarioid":         t.UsuarioID,
		"usuariorelacionid": t.UsuarioRelacionID,
	}

	var resultado models.Relacion //el find va a querer crear un doc por eso guardamos el el resultado en resultado
	fmt.Println(resultado)
	err := col.FindOne(ctx, condicion).Decode(&resultado) // si FindOne tiene un result hacemos un decode a resultado
	if err != nil {
		fmt.Println(err.Error())
		return false, err
	}
	return true, nil
}
