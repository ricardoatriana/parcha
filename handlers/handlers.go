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
	router.HandleFunc("/verperfil", middlew.ChequeoBD(middlew.ValidoJWT(routers.VerPerfil))).Methods("GET")
	router.HandleFunc("/modificarPerfil", middlew.ChequeoBD(middlew.ValidoJWT(routers.ModificarPerfil))).Methods("PUT")
	router.HandleFunc("/parla", middlew.ChequeoBD(middlew.ValidoJWT(routers.GraboParla))).Methods("POST")
	router.HandleFunc("/leoParlas", middlew.ChequeoBD(middlew.ValidoJWT(routers.LeoParlas))).Methods("GET")
	router.HandleFunc("/eliminarParla", middlew.ChequeoBD(middlew.ValidoJWT(routers.EliminarParla))).Methods("DELETE")

	router.HandleFunc("/subirAvatar", middlew.ChequeoBD(middlew.ValidoJWT(routers.SubirAvatar))).Methods("POST")
	router.HandleFunc("/obtenerAvatar", middlew.ChequeoBD(routers.ObtenerAvatar)).Methods("GET") // no lleva el chequeo de TOken, puede ser vistos in logearse a la app
	router.HandleFunc("/subirBanner", middlew.ChequeoBD(middlew.ValidoJWT(routers.SubirBanner))).Methods("POST")
	router.HandleFunc("/obtenerBanner", middlew.ChequeoBD(routers.ObtenerBanner)).Methods("GET") // no lleva el chequeo de TOken, puede ser vistos in logearse a la app

	router.HandleFunc("/altaRelacion", middlew.ChequeoBD(middlew.ValidoJWT(routers.AltaRelacion))).Methods("POST")
	router.HandleFunc("/bajaRelacion", middlew.ChequeoBD(middlew.ValidoJWT(routers.BajaRelacion))).Methods("DELETE")
	router.HandleFunc("/consultaRelacion", middlew.ChequeoBD(middlew.ValidoJWT(routers.ConsultaRelacion))).Methods("GET")

	router.HandleFunc("/listaUsuarios", middlew.ChequeoBD(middlew.ValidoJWT(routers.ListaUsuarios))).Methods("GET")
	router.HandleFunc("/leoParlasSeguidores", middlew.ChequeoBD(middlew.ValidoJWT(routers.LeoParlasRelacion))).Methods("GET")

	PORT := os.Getenv("PORT")
	if PORT == "" {
		PORT = "8080"
	}
	handler := cors.AllowAll().Handler(router) //El router se le pasa a cors, y mira los permisos el cors. Permisos a determinadas direccionesIP, para q dependiendo del llamado a la IP pueda hace determinadas funciones
	log.Fatal(http.ListenAndServe(":"+PORT, handler))
}
