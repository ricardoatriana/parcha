package routers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/ricardoatriana/parcha/bd"
)

/*ListaUsuarios leo la lista de los usuarios. Capturo los parametros, hago una lectura de LEOUSUARIOSTODODS luego si andubo todo bien decodifico el json*/
func ListaUsuarios(w http.ResponseWriter, r *http.Request) {

	typeUser := r.URL.Query().Get("type") //capturar el parametro type, son 2 tipos de usuario los q sigo y los q no sigo
	page := r.URL.Query().Get("page")
	search := r.URL.Query().Get("search")

	pagTemp, err := strconv.Atoi(page) //atoi cambia el alfabetico en entero, se le llama temp porq luego se va a convertir en entero e 64
	if err != nil {
		http.Error(w, "Debe enviar el parametro pagina como entero mayor a 0", http.StatusBadRequest)
		return
	}

	pag := int64(pagTemp)

	result, status := bd.LeoUsuariosTodos(IDUsuario, pag, search, typeUser)
	if status == false {
		http.Error(w, "Error al leer los usuarios", http.StatusBadRequest)
		return	
	}
	w.Header().Set("Content-type", "application/json") // se setea de tipo json y enviamos los valores
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(result)
}
