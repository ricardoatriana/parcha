package bd

import (
	"github.com/ricardoatriana/parcha/models"
	"golang.org/x/crypto/bcrypt"
)

/*IntentoLogin realiza el chequeo de Login a la bd */
func IntentoLogin(email string, clave string) (models.Usuario, bool) {
	usu, encontrado, _ := ChequeoYaExisteUsuario(email) //trae el usuario y el bool encontrado. Como no se usa
	if encontrado == false {
		return usu, false
	}

	claveBytes := []byte(clave)  // el bcrypt trabaja con slice byte osea []byte. Esta clave no esta encriptada
	claveBD := []byte(usu.Clave) // esta clave si esta encriptada
	err := bcrypt.CompareHashAndPassword(claveBD, claveBytes)
	if err != nil {
		return usu, false
	}
	return usu, true
}
