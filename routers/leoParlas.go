package routers

import (
	"encoding/json"
	"net/http"
	"strconv" //para hacer conversion de datos

	"github.com/ricardoatriana/parcha/bd"
)

/*LeoParlas lee las parlas*/
func LeoParlas(w http.ResponseWriter, r *http.Request) {
	ID := r.URL.Query().Get("id")
	if len(ID) < 1 {
		http.Error(w, "Debe enviar el parametro ID", http.StatusBadRequest)
		return
	}

	if len(r.URL.Query().Get("pagina")) < 1 { //pagina es un string
		http.Error(w, "Debe enviar el parametro pagina", http.StatusBadRequest)
		return
	}
	pagina, err := strconv.Atoi(r.URL.Query().Get("pagina")) //pagina se convierte a un entero
	if err != nil {
		http.Error(w, "Debe enviar el parametro pagina con un valor mayor a 0", http.StatusBadRequest)
		return
	}
	pag := int64(pagina)                         //convertimos pagina de int a int64, porq la rutina q usamos para paginar en bson reuqiere q ese dato sea en int64
	respuesta, correcto := bd.LeoParlas(ID, pag) //esta func devolvia un modelo y booleano q esta en correcto o no
	if correcto == false {
		http.Error(w, "Error al leer las parlas", http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-type", "application/json") // se setea de tipo json y enviamos los valores
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(respuesta)
}
