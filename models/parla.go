package models 

/*Parla captura del Body, el msj q nos llega. Este modelo es para decodifcar el body q viene con Json*/
type Parla struct { //esta en mayus para poder exportarlo
	Mensaje string `bson:"mensaje" json:"mensaje"`
}