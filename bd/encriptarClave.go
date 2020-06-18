package bd

import "golang.org/x/crypto/bcrypt" //se usa tbn en nodejs

/*EncriptarClave es la rutina q me permite encriptar clave*/
func EncriptarClave(clave string) (string, error) {
	costo := 8
	bytes, err := bcrypt.GenerateFromPassword([]byte(clave), costo)
	return string(bytes), err
}
