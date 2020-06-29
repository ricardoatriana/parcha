package bd

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

/*BorroParla borra una parla determinada*/
func BorroParla(ID string, UserID string) error { //El 1ro es el id del tweet y el otro es el id del usuario
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*15)
	defer cancel()

	db := MongoCN.Database("parcha")
	col := db.Collection("parla")

	objID, _ := primitive.ObjectIDFromHex(ID) //se crea porq el ID de la parla hay q convertirlo en un ObjetID
	condicion := bson.M{
		"_id":    objID,
		"userid": UserID,
	}
	_, err := col.DeleteOne(ctx, condicion)
	return err
}
