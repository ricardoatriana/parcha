package routers

import (
	"io" //inputoutput muy importante
	"net/http"
	"os"
	"strings"

	"github.com/ricardoatriana/parcha/bd"
	"github.com/ricardoatriana/parcha/models"
)

/*SubirBanner sube el avatar al servidor*/
func SubirBanner(w http.ResponseWriter, r *http.Request) { //vamos a capturar un archivo desde el request porq en postman vamos a subir un archivo y tbn lo va a capturar como si fuera el front
	file, handler, err := r.FormFile("banner")
	var extension = strings.Split(handler.Filename, ".")[1]               //vamos a capturar del arch q nos vino la extension. Lo dividimos en el punto, cuando se divide se convierte en un vector y capturamos el 1er elemento del vector, ya convertido en string porque al final se coloca [1]
	var archivo string = "uploads/banners/" + IDUsuario + "." + extension //IDUsuario es un string. Se hace esto porque el usuario envia un avatar q se toma desde la pc del susuario q se llama "mifoto", y asi no se puede grabar porq habria colisiones enre usuarios q teng amisma foto

	f, err := os.OpenFile(archivo, os.O_WRONLY|os.O_CREATE, 0666) //se crea una func del OS q devuelve 2 parametros q abre el arch. 0666 son los permisos de lectura,escritura y ejecucion. Los atributos son WriteONLY y Create. Practicamente esta creando el archivo, porq se ha reservado un espacio en el disco para saber si se creo o no
	if err != nil {
		http.Error(w, "Error al subir la imagen !"+err.Error(), http.StatusBadRequest)
		return
	}

	_, err = io.Copy(f, file) //va a copiar el arch en f. Con el copy lo q se logra es q se grabe en disco
	if err != nil {
		http.Error(w, "Error al copiar la imagen !"+err.Error(), http.StatusBadRequest)
		return
	}
	//una vez copiado el arch, vamos a la BD a modificar el registro del usuario y le dicimos q el Avatar va a estar en el lugar q va a estar
	var usuario models.Usuario
	var status bool
	usuario.Banner = IDUsuario + "." + extension
	status, err = bd.ModificoRegistro(usuario, IDUsuario)
	if err != nil || status == false { // || esto sifnifica o en golang
		http.Error(w, "Error al grabar banner en la BD !"+err.Error(), http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-type", "application/json")
	w.WriteHeader(http.StatusCreated)

}
