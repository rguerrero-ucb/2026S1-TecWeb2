package infrastructure

import (
	"database/sql"
	"log/slog"
	"ucb/todo/internal/service"
)

type TareaRepositoryImpl struct {
	Db *sql.DB
}

func NewTareaRepository(db *sql.DB) *TareaRepositoryImpl {
	return &TareaRepositoryImpl{Db: db}
}

func (r *TareaRepositoryImpl) FindAll() ([]service.Tarea, error) {
	rows, err := r.Db.Query("SELECT id, nombre, estado FROM todo.tareas")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var tareas []service.Tarea
	for rows.Next() {
		var tarea service.Tarea
		err := rows.Scan(&tarea.ID, &tarea.Nombre, &tarea.Estado)
		if err != nil {
			return nil, err
		}
		tareas = append(tareas, tarea)
	}
	return tareas, nil
}

func (r *TareaRepositoryImpl) NextIdentity() (int, error) {
	var id int
	err := r.Db.QueryRow("SELECT nextval('todo.tareas_id_seq')").Scan(&id)
	if err != nil {
		return 0, err
	}
	return id, nil
}

func (r *TareaRepositoryImpl) Create(tarea service.Tarea) error {
	slog.Debug("Guardando tarea en repository", "Tarea", tarea)
	_, err := r.Db.Exec("INSERT INTO todo.tareas (id, nombre, estado) VALUES ($1, $2, $3)",
		tarea.ID,
		tarea.Nombre,
		tarea.Estado,
	)
	if err != nil {
		return err
	}
	return nil
}

func (r *TareaRepositoryImpl) Delete(id int) error {
	_, err := r.Db.Exec("DELETE FROM todo.tareas WHERE id = $1", id)
	return err
}
