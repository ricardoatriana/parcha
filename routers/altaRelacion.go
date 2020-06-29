package routers

import (
	"net/http"

	"github.com/ricardoatriana/parcha/bd"
	"github.com/ricardoatriana/parcha/models"
)

/*AltaRelacion realiza el registro de la relacion entre usuarios*/
func AltaRelacion(w http.ResponseWriter, r *http.Request) {
	ID := r.URL.Query().Get("id") // de mi url capturar el parametro id
	if len(ID) < 1 {
		http.Error(w, "Debe enviar el parametro ID", http.StatusBadRequest)
		return
	}

	var t models.Relacion
	t.UsuarioID = IDUsuario  //UsuarioID soy yo y lo saco de IDUsuario
	t.UsuarioRelacionID = ID //q usuario voy a seguir

	status, err := bd.InsertoRelacion(t) // se le pasa el modelo linea17
	if err != nil {
		http.Error(w, "Ocurrio un error al insertar relacion"+err.Error(), http.StatusBadRequest)
		return
	}
	if status == false {
		http.Error(w, "No se ha logrado insertar la relacion"+err.Error(), http.StatusBadRequest)
		return
	}
	w.WriteHeader(http.StatusCreated)

}
