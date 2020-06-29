package models

/*Relacion modelo para grabar la relacion de un usuario con otro*/
type Relacion struct {
	UsuarioID         string `bson:"usuarioid" json:"usuarioID"`                 //mi ID
	UsuarioRelacionID string `bson:"usuariorelacionid" json:"usuarioRelacionId"` //ID del usuario q estoy siguiendo
}
