package service

type Tarea struct {
	ID     int    `json:"id"`
	Nombre string `json:"nombre"`
	Estado int8   `json:"estado"`
}
