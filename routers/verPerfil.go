package routers

import (
	"encoding/json"
	"net/http"

	"github.com/ricardoatriana/parcha/bd"
)

/*VerPerfil permite extraer los valores del perfil*/
func VerPerfil(w http.ResponseWriter, r *http.Request) { //Todos nuestros endpoints no devuelven nada orqu son metodos. Ahora vamos a extraer de nuestro Body los parametros q vinieron
	ID := r.URL.Query().Get("id") //extraemos de la URL el parametro ID,
	if len(ID) < 1 {              //fue a buscar esta var y no la encontro
		http.Error(w, "Debe enviar el parametro ID", http.StatusBadRequest)
		return
	}

	perfil, err := bd.BuscoPerfil(ID) //pero si todo anduvo bien se crea una var de tipo models Usuario. Se crea automaticamente porque en bd BuscoPerfil lo q devuelve es un modelsUsuario
	if err != nil {
		http.Error(w, "Ocurrio un error al intentar buscar el registro "+err.Error(), 400)
		return
	}
	w.Header().Set("context-type", "application/json") //si todo esta bien se le avisa q lo q esta enviando es un Json
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(perfil)
}
