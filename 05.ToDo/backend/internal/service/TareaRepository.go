package service

type TareaRepository interface {
	NextIdentity() (int, error)
	FindAll() ([]Tarea, error)
	Create(tarea Tarea) error
	Delete(id int) error
}
