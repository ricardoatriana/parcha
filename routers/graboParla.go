package routers

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/ricardoatriana/parcha/bd"
	"github.com/ricardoatriana/parcha/models"
)

/*GraboParla permite grabar el parla en la BD*/
func GraboParla(w http.ResponseWriter, r *http.Request) { //recibimos el json en nuestro Body y lo hemos decodifcado dentro de mensaje
	var mensaje models.Parla
	err := json.NewDecoder(r.Body).Decode(&mensaje)

	registro := models.GraboParla{
		UserID:  IDUsuario,       //esta var IDUsuario es global
		Mensaje: mensaje.Mensaje, //el mensaje q acabamos de decodificar y adentro tiene un campo llamado Mensaje q es el q nos viene en el json del Body
		Fecha:   time.Now(),
	}

	_, status, err := bd.InsertoParla(registro)
	if err != nil {
		http.Error(w, "Ocurrio un error al intentar insertar el registro, reintente nuevamente "+err.Error(), 400)
		return
	}
	if status == false {
		http.Error(w, "No se ha logrado insertar el tweet", 400)
		return
	}
	w.WriteHeader(http.StatusCreated)
}
