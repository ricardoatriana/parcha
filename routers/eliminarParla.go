package routers

import (
	"net/http"

	"github.com/ricardoatriana/parcha/bd"
)

/*EliminarParla podria llamarse tbn borrarparla pero se escribe asi para no confudirnos*/
func EliminarParla(w http.ResponseWriter, r *http.Request) {
	ID := r.URL.Query().Get("id")
	if len(ID) < 1 {
		http.Error(w, "Debe enviar el parametro ID", http.StatusBadRequest)
		return
	}

	err := bd.BorroParla(ID, IDUsuario) //el 1ro es el ID de la URL y el 2do es el IDUusuario q venia en mi token
	if err != nil {
		http.Error(w, "Ocurrio un error al intentar borrar el ID"+err.Error(), http.StatusBadRequest)
		return
	}
	w.Header().Set("Content-type", "application/json") // si no se escribe esta linea el navegador no sabe en q formato se le esta enviando
	w.WriteHeader(http.StatusCreated)
}
