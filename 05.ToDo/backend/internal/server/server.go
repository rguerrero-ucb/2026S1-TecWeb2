package server

import (
	"encoding/json"
	"fmt"
	"log/slog"
	"net/http"
	"os"
	"strconv"
	"ucb/todo/internal/repository"
	"ucb/todo/internal/service"
)

func SendResponse(w http.ResponseWriter, response map[string]any) {
	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(response); err != nil {
		http.Error(w, "Error al codificar la respuesta", http.StatusInternalServerError)
	}
}

// Middleware CORS para permitir requests desde cualquier origen
func corsMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")

		// Manejar preflight requests
		if r.Method == http.MethodOptions {
			w.WriteHeader(http.StatusOK)
			return
		}

		next.ServeHTTP(w, r)
	})
}

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

	// Inicializando repositorios
	tareaRepository := repository.NewTareaRepository(db)

	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "¡Hola, mundo!")
	})

	mux.HandleFunc("/tareas", func(w http.ResponseWriter, r *http.Request) {
		tareaService := service.NewTareaService(tareaRepository)
		tareas, err := tareaService.GetAllTareas()
		if err != nil {
			http.Error(w, "Error al obtener las tareas", http.StatusInternalServerError)
			return
		}

		SendResponse(w, map[string]any{
			"success": true,
			"tareas":  tareas,
		})
	})

	mux.HandleFunc("/tareas/create", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			http.Error(w, "Método no permitido", http.StatusMethodNotAllowed)
			return
		}

		nombre := r.FormValue("nombre")
		if nombre == "" {
			SendResponse(w, map[string]any{
				"success": false,
				"error":   "El campo 'nombre' es requerido",
			})
			return
		}

		tareaService := service.NewTareaService(tareaRepository)
		tarea, err := tareaService.CreateTarea(nombre)
		if err != nil {
			SendResponse(w, map[string]any{
				"success": false,
				"error":   "Error al crear la tarea",
			})
			return
		}
		SendResponse(w, map[string]any{
			"success": true,
			"tarea":   tarea,
		})

	})

	mux.HandleFunc("/tareas/delete", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			http.Error(w, "Método no permitido", http.StatusMethodNotAllowed)
			return
		}

		paramId := r.FormValue("id")
		if paramId == "" {
			SendResponse(w, map[string]any{
				"success": false,
				"error":   "El campo 'id' es requerido.",
			})
			return
		}

		id, err := strconv.Atoi(paramId)
		if err != nil {
			SendResponse(w, map[string]any{
				"success": false,
				"error":   "El campo 'id' debe ser un número válido",
			})
			return
		}

		tareaService := service.NewTareaService(tareaRepository)
		err = tareaService.DeleteTarea(id)
		if err != nil {
			SendResponse(w, map[string]any{
				"success": false,
				"error":   "Error al eliminar la tarea: " + err.Error(),
			})
			return
		}

		SendResponse(w, map[string]any{
			"success": true,
		})
	})

	fmt.Printf("Servidor escuchando en http://localhost:8080\n")
	http.ListenAndServe(":8080", corsMiddleware(mux))
}
