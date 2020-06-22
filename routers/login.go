package routers

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/ricardoatriana/parcha/bd"
	"github.com/ricardoatriana/parcha/jwt"
	"github.com/ricardoatriana/parcha/models"
)

/*Login realiza el logeo*/
func Login(w http.ResponseWriter, r *http.Request) { //No va a devolver nada porque es un endpoint. TOdo lo q va en router son endpoints
	w.Header().Add("content-type", "application/json") //seteamos el header de nuestro objeto ResponseWriter, esta es la respuesta. Entonces en nuestro header de la respuesta tenemos q iicar q lo q devolvemos es en fomrato json

	var t models.Usuario

	err := json.NewDecoder(r.Body).Decode(&t)
	if err != nil {
		http.Error(w, "Usuario y/o Contraseña inválida"+err.Error(), 400)
		return
	}
	if len(t.Email) == 0 {
		http.Error(w, "El email del usuario es requerido", 400)
		return
	}

	documento, existe := bd.IntentoLogin(t.Email, t.Clave) //IntentoLogin es false es q no pudo lograrlo
	if existe == false {
		http.Error(w, "Usuario y/o Contraseña inválida", 400)
		return
	}

	jwtKey, err := jwt.GeneroJWT(documento) //Esta func GeneroJWT recibe el documento y devuelve el token en modo string para q se pueda devolver al usuario con el http
	if err != nil {
		http.Error(w, "Ocurrio un error al intentar generar un token correspondiente "+err.Error(), 400)
		return
	}

	resp := models.RespuestaLogin{
		Token: jwtKey,
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(resp) //Se encodigfica la respuesta, respuesta tenia un json donde venia el token

	//como se graba desde el back una cookie
	expirationTime := time.Now().Add(24 * time.Hour)
	http.SetCookie(w, &http.Cookie{
		Name:    "token",
		Value:   jwtKey,
		Expires: expirationTime,
	})
}
