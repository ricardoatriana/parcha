package middlew

import (
	"net/http"

	"github.com/ricardoatriana/parcha/bd" //desde aquise va a chequear la bd de datos si esta online, pero tiene q traer las conexciones q estan en el paquete BD
)

//ChequeoBD es el middlew q me permite conocer el estado de la BD
func ChequeoBD(next http.HandlerFunc) http.HandlerFunc { //lo que va a recibir es lo q recibe todos los HandleFunc
	return func(w http.ResponseWriter, r *http.Request) {
		if bd.ChequeoConnection() == 0 {
			http.Error(w, "Conexion perdida con la BD", 500)
			return
		}
		next.ServeHTTP(w, r) //Si no dio error pasa todos los objetos de writer y request
	}

}
