package service

import "log/slog"

type TareaService struct {
	tareaRepository TareaRepository
}

func NewTareaService(tareaRepository TareaRepository) *TareaService {
	return &TareaService{tareaRepository: tareaRepository}
}

func (s *TareaService) GetAllTareas() ([]Tarea, error) {
	return s.tareaRepository.FindAll()
}

func (s *TareaService) CreateTarea(nombre string) (Tarea, error) {
	slog.Debug("Creando tarea en service", "Nombre", nombre)
	ID, err := s.tareaRepository.NextIdentity()
	if err != nil {
		return Tarea{}, err
	}
	tarea := Tarea{
		ID:     ID,
		Nombre: nombre,
		Estado: 0,
	}
	slog.Debug("Guardando tarea en service", "Tarea", tarea)
	err = s.tareaRepository.Create(tarea)
	if err != nil {
		return Tarea{}, err
	}
	slog.Debug("Tarea creada en service", "Tarea", tarea)
	return tarea, nil
}

func (s *TareaService) DeleteTarea(id int) error {
	return s.tareaRepository.Delete(id)
}
