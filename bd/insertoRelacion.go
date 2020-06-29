package bd

import (
	"context"
	"time"

	"github.com/ricardoatriana/parcha/models"
)

/*InsertoRelacion graba la relacion en a BD*/
func InsertoRelacion(t models.Relacion) (bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*15)
	defer cancel()

	db := MongoCN.Database("parcha")
	col := db.Collection("relacion")

	_, err := col.InsertOne(ctx, t) //se inserta el modelo en la BD
	if err != nil {
		return false, err
	}

	return true, nil
}
