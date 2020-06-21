package jwt

import (
	"time"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/ricardoatriana/parcha/models"
)

/*GeneroJWT genera el encriptado con JWT*/
func GeneroJWT(t models.Usuario) (string, error) { //el string tiene el token y un error

	miClave := []byte("RicardoElGuapo")

	payload := jwt.MapClaims{ //En MapClaims grabamos la lista de privilegios
		"email":            t.Email,
		"nombre":           t.Nombre,
		"apellido":         t.Apellido,
		"fecha_nacimiento": t.FechaNacimiento,
		"biografia":        t.Biografia,
		"ubicacion":        t.Ubicacion,
		"sitio web":        t.SitioWeb,
		"_id":              t.ID.Hex(),
		"exp":              time.Now().Add(time.Hour * 24).Unix(), //El unix es un formato muy rapido, se graba muy fino y liviano
	}

	token := jwt.NewWithClaims(jwt.SigningMethodES256, payload) //Estamos en el Header de JWT, estamos eligiendo el algoritmo q tiene elegir para encriptar la clave
	tokenStr, err := token.SignedString(miClave)
	if err != nil {
		return tokenStr, err //Tokenstr va a estar vacio porque hay un error
	}
	return tokenStr, nil
}
