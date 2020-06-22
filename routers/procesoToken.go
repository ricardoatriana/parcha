package routers

import (
	"errors"
	"strings"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/ricardoatriana/parcha/bd"
	"github.com/ricardoatriana/parcha/models"
)

/*Email valor de email usado en todos los EndPoints*/
var Email string

/*IDUsuario es el ID devuelto del modelo que se usara en todos los EndPoints*/
var IDUsuario string

/*ProcesoToken para extraer los valores del token*/
func ProcesoToken(tk string) (*models.Claim, bool, string, error) { //El error hay q ponerlo al final
	miPwrd := []byte("RicardoElGuapo")
	claims := &models.Claim{}

	splitToken := strings.Split(tk, "Bearer") //Esta funcion divide el texto de tk del delimitador q es Bearer. Splittoken se convierte en un vector en el cual el elemento 0 se tiene Bearer y en el 1 el token.
	if len(splitToken) != 2 {
		return claims, false, string(""), errors.New("formato de token invalido") //El string vacio ahi va el ID
	}

	tk = strings.TrimSpace(splitToken[1]) //esta funcion le quita los espacios , es como quitarle la palabra bearer

	tkn, err := jwt.ParseWithClaims(tk, claims, func(token *jwt.Token) (interface{}, error) {
		return miPwrd, nil
	})
	if err == nil {
		_, encontrado, _ := bd.ChequeoYaExisteUsuario(claims.Email)
		if encontrado == true {
			Email = claims.Email //no se le pone dos puntos porque ya fue creada
			IDUsuario = claims.ID.Hex()
		}
		return claims, encontrado, IDUsuario, nil
	}
	if !tkn.Valid {
		return claims, false, string(""), errors.New("token Invalido")
	}
	return claims, false, string(""), err
}
