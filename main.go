package main

import (
	"log"

	"github.com/ricardoatriana/parcha/bd"
	"github.com/ricardoatriana/parcha/handlers"
)

func main() {
	if bd.ChequeoConnection() == 0 { // ChequeoConnection hace un Ping a la BD, recorda q escribimos q devolvia un 1 o 0
		log.Fatal("Sin conexion a la BD")
		return
	}
	handlers.Manejadores()
}
