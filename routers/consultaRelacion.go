package routers

import (
	"encoding/json"
	"net/http"

	"github.com/ricardoatriana/parcha/bd"
	"github.com/ricardoatriana/parcha/models"
)

/*ConsultaRelacion chequea si hay relacion entre 2 usuarios*/
func ConsultaRelacion(w http.ResponseWriter, r *http.Request) {
	ID := r.URL.Query().Get("id") // de mi url capturar el parametro id

	var t models.Relacion
	t.UsuarioID = IDUsuario  //UsuarioID soy yo y lo saco de IDUsuario
	t.UsuarioRelacionID = ID //q usuario voy a seguir

	var resp models.RespuestaConsultaRelacion //luego consultamos rutina BD consultoRelacion
	status, err := bd.ConsultoRelacion(t)
	if err != nil || status == false { // aqui no hay mostrar error, sino comunicarla directamente a la maquina a traves de http
		resp.Status = false
	} else {
		resp.Status = true
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(resp) //econdificamos porque ya tenemos nuestro modelo relacion y nosotros lo encodificamos para enviar a nuestro navegdor
}
