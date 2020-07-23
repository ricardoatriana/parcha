package routers

import (
	"io" //inputoutput muy importante
	"net/http"
	"os"

	"github.com/ricardoatriana/parcha/bd"
)

/*ObtenerBanner sube el avatar al servidor. Resumen se pasa el ID q me vino, lo chequeo, hago una busqueda de perfil por ese ID, lo cargo en un modelo, sino hubo un error sigo, intento abrir el archivo sino hubo un error sigo y hago una copia del arch a mi response writer */
func ObtenerBanner(w http.ResponseWriter, r *http.Request) {

	ID := r.URL.Query().Get("id") //De la URL vamos vamos a tener q enviar el id de quien estamos obteniendo el avatar. Porq necesitamos obtner de la app no solo el avatar mio sino de cada usuario
	if len(ID) < 1 {
		http.Error(w, "Debe enviar el parametro ID", http.StatusBadRequest)
		return
	}
	//Ahora buscamos un perifl del usuario porq vamos a traer en abase a ese id el perfil del usuario
	perfil, err := bd.BuscoPerfil(ID)
	if err != nil {
		http.Error(w, "Usuario no encontrado", http.StatusBadRequest)
		return
	}
	//Ahora eleemos el avatar del usuario
	OpenFile, err := os.Open("uploads/banners/" + perfil.Banner)
	if err != nil {
		http.Error(w, "Imagen no encontrado", http.StatusBadRequest)
		return
	}

	//Envio de la imagen al http. Se hace una copia del archivo a mi responsewriter
	_, err = io.Copy(w, OpenFile) //La imagen es enviada al ResponseWriter. Este no tiene prueba en postman
	if err != nil {
		http.Error(w, "Error al copiar la imagen", http.StatusBadRequest) //El fr
	}
}
