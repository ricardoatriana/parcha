package routers

import (
	"encoding/json"
	"net/http"

	"github.com/ricardoatriana/parcha/bd"
	"github.com/ricardoatriana/parcha/models"
)

/*ModificarPerfil modifica el perfil del usuario*/
func ModificarPerfil(w http.ResponseWriter, r *http.Request) {

	var t models.Usuario //creo un modelo de usuario grabo el body y se decodifica en este modelo de usuario

	err := json.NewDecoder(r.Body).Decode(&t)
	if err != nil {
		http.Error(w, "Datos Incorrectos "+err.Error(), 400) //El body vino vacio o errores de formato, un json mal construido
		return
	}

	var status bool //Se le asigna a la var y al objeto error el resultado de modifico registro donde va el modelo de usuario q acaba de codifcar del body y ademas la var Global q se ha venido capturado con el middlew, se envian esos 2 datos

	status, err = bd.ModificoRegistro(t, IDUsuario) //modifico registro devolvia un Bool y un err, entonces el bool se graba en status. Si no se llama la var Global, porje plo en ModificarPerfil tendria q ejecutar un llamado a procso token para decodificar el token q recibo y ahi extrar el ID de usuario
	if err != nil {
		http.Error(w, "Ocurrio un error a intentar modificar el registro. Reintente nuevamente"+err.Error(), 400)
		return
	}
	if status == false {
		http.Error(w, "No se ha logrado modificar el registro del usuario", 400)
		return
	}
	w.WriteHeader(http.StatusCreated)
}
