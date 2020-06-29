package bd

import (
	"context"
	"time"

	"github.com/ricardoatriana/parcha/models"
)

/*BorroRelacion borro relacion en la BD*/
func BorroRelacion(t models.Relacion) (bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*15)
	defer cancel()

	db := MongoCN.Database("parcha")
	col := db.Collection("relacion")

	_, err := col.DeleteOne(ctx, t)
	if err != nil {
		return false, err //booleano primero false y luego el erro
	}
	return true, nil
}
