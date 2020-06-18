package routers

import (
	"encoding/json" //paquete manipulacion de datos json
	"net/http"

	"github.com/ricardoatriana/parcha/bd"
	"github.com/ricardoatriana/parcha/models"
)

/*Registro es la func para crear en la BD el registro de usuario*/
func Registro(w http.ResponseWriter, r *http.Request) { //metodo porque no deuvelve valores o parametros
	//vamos a llamar a nuestro modelo
	var t models.Usuario
	//crea un modelo json de usuario luego el body lo decodifica en el modelo t
	err := json.NewDecoder(r.Body).Decode(&t) //el body de un http es un stream, quiere decir q es un dato q se puede leer solo 1 vez, entonces no se puede usar el body en varios lugares
	if err != nil {
		http.Error(w, "Error en los datos recibidos"+err.Error(), 400) //el err.error() complemente el texto q se devuelve en español con el error q revota de json
		return                                                         // el return acaba con todo
	}
	if len(t.Email) == 0 {
		http.Error(w, "El email de usuario es requerido", 400) //el err.error() porque ya sabemos cual fue el error
		return                                                 // el return acaba con todo
	}
	if len(t.Clave) < 6 {
		http.Error(w, "No cumple con el minimo de longitud para clave", 400) //el err.error() porque ya sabemos cual fue el error
		return                                                               // el return acaba con todo
	}

	_, encontrado, _ := bd.ChequeoYaExisteUsuario(t.Email) //para evitar carga la misma info de registro
	if encontrado == true {
		http.Error(w, "Ya existe un usuario registrado con ese email", 400) //el err.error() porque ya sabemos cual fue el error
		return                                                              // el return acaba con todo
	}
	_, status, err := bd.InsertoRegistro(t)
	if err != nil {
		http.Error(w, "Ocurrio un error al intentar realizar el registro de usuario"+err.Error(), 400) //el err.error() porque ya sabemos cual fue el error
		return                                                                                         // el return acaba con todo
	}
	if status == false { //le dio un status vacio
		http.Error(w, "No se ha logrado inserta el registro de usuario"+err.Error(), 400) //el err.error() porque ya sabemos cual fue el error
		return
	}
	w.WriteHeader(http.StatusCreated)
}
