package middlew

import (
	"net/http"

	"github.com/ricardoatriana/parcha/routers"
)

/*ValidoJWT permite validar el JWT que nos viene en la peticion*/
func ValidoJWT(next http.HandlerFunc) http.HandlerFunc { //como todo middlew va a ser http handle func
	return func(w http.ResponseWriter, r *http.Request) {
		_, _, _, err := routers.ProcesoToken(r.Header.Get("Authorization")) //Es una rutina q le vamos a pasar el token si es valido o no. Autho es la var que se va apsaar a tring y lo pasamos a parametroa proceso token
		if err != nil {
			http.Error(w, "Error en el token !"+err.Error(), http.StatusBadRequest)
			return
		}
		next.ServeHTTP(w, r)
	}

}
