package server

import (
	"encoding/json"
	"fmt"
	"log/slog"
	"net/http"
	"os"
	"ucb/todo/internal/repository"
	"ucb/todo/internal/service"
)

func Run() {
	logger := slog.New(slog.NewTextHandler(os.Stdout, nil))
	slog.SetDefault(logger)

	mux := http.NewServeMux()

	db := repository.InitDB()
	if db == nil {
		fmt.Println("No se pudo establecer la conexión a la base de datos. El servidor no se iniciará.")
		return
	}
	defer db.Close()

	tareaRepository := repository.NewTareaRepository(db)

	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "¡Hola, mundo!")
	})

	mux.HandleFunc("/tareas", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodGet {
			tareaService := service.NewTareaService(tareaRepository)
			tareas, err := tareaService.GetAllTareas()
			if err != nil {
				http.Error(w, "Error al obtener las tareas", http.StatusInternalServerError)
				return
			}
			w.Header().Set("Content-Type", "application/json")
			if err := json.NewEncoder(w).Encode(tareas); err != nil {
				http.Error(w, "Error al codificar las tareas", http.StatusInternalServerError)
			}
		}
	})

	mux.HandleFunc("/tareas/create", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodPost {
			tareaService := service.NewTareaService(tareaRepository)
			nombre := r.FormValue("nombre")
			if nombre == "" {
				http.Error(w, "El campo 'nombre' es requerido", http.StatusBadRequest)
				return
			}
			tarea, err := tareaService.CreateTarea(nombre)
			if err != nil {
				http.Error(w, "Error al crear la tarea", http.StatusInternalServerError)
				return
			}
			w.Header().Set("Content-Type", "application/json")
			if err := json.NewEncoder(w).Encode(tarea); err != nil {
				http.Error(w, "Error al codificar la tarea", http.StatusInternalServerError)
			}
		}
	})

	fmt.Printf("Servidor escuchando en http://localhost:8080\n")
	http.ListenAndServe(":8080", mux)
}
