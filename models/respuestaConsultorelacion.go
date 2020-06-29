package models

/*RespuestaConsultaRelacion tiene el true o false que obtiene de consultar la relacion*/
type RespuestaConsultaRelacion struct {
	Status bool `json:"status"` // no necesito q tenga bson ya que no es para grabar en la base de datos, sino es para devolver en el http
}
