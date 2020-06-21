package handlers

import (
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/ricardoatriana/parcha/middlew"
	"github.com/ricardoatriana/parcha/routers"

	"github.com/rs/cors" //son los permisos dados a la API para q sea accesible desde caulquier lugar
)

//Manejadores seteo mi puerto, el Handler y porngo a escuchar el servidor. Se van a alistar las 16 rutas que se van a manejar
func Manejadores() {
	router := mux.NewRouter() //se crea un objeto llamado router. Lo q hace el mux es capturar la http y va a manejar el ResponseWriter y la request q viene de la API. Y va a ver si en el llamado hay info, en el header y va a enviar la respuesta al navegador
	//la API va a devolver un status si se pudo logear en un token, deolviendo info al mismo tiempo q recibe
	//aqui vendrian todas las rutas a manejar, pero ahora se abre el puerto
	router.HandleFunc("/registro", middlew.ChequeoBD(routers.Registro)).Methods("POST") //cuando en el navegador se le pone un /registro lo procesa a traves del llamaod POST ejecutando el middleware y este al BD si nohay problema devuelve el control al router
	router.HandleFunc("/login", middlew.ChequeoBD(routers.Login)).Methods("POST")

	PORT := os.Getenv("PORT")
	if PORT == "" {
		PORT = "8080"
	}
	handler := cors.AllowAll().Handler(router) //El router se le pasa a cors, y mira los permisos el cors. Permisos a determinadas direccionesIP, para q dependiendo del llamado a la IP pueda hace determinadas funciones
	log.Fatal(http.ListenAndServe(":"+PORT, handler))
}
