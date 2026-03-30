package server

import (
	"encoding/json"
	"fmt"
	"log/slog"
	"net/http"
	"os"
	"ucb/todo/internal/infrastructure"
	"ucb/todo/internal/service"
)

func SendResponse(w http.ResponseWriter, response map[string]any) {
	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(response); err != nil {
		http.Error(w, "Error al codificar la respuesta", http.StatusInternalServerError)
	}
}

// CORSMiddleware añade headers CORS a todas las respuestas
func CORSMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Permitir solicitudes desde cualquier origen
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type, ngrok-skip-browser-warning")
		w.Header().Set("Access-Control-Max-Age", "86400")

		// Manejar solicitudes preflight (OPTIONS)
		if r.Method == http.MethodOptions {
			w.WriteHeader(http.StatusOK)
			return
		}

		next.ServeHTTP(w, r)
	})
}

func Run() {
	// Configurando el logger
	logger := slog.New(slog.NewTextHandler(os.Stdout, nil))
	slog.SetDefault(logger)

	// Inicializando la base de datos
	db := infrastructure.InitDB()
	if db == nil {
		fmt.Println("No se pudo establecer la conexión a la base de datos. El servidor no se iniciará.")
		return
	}
	defer db.Close()

	// Inicializando repositorios
	tareaRepository := infrastructure.NewTareaRepository(db)
	tareaService := service.NewTareaService(tareaRepository)

	// Inicializando handlers
	handler := NewHandler(tareaService)

	// Registrando rutas
	mux := http.NewServeMux()
	mux.HandleFunc("/", handler.HomeHandler)
	mux.HandleFunc("/tareas", handler.TareasHandler)
	// Mantener rutas antiguas para compatibilidad
	mux.HandleFunc("/tareas/create", handler.CreateTareaHandler)
	mux.HandleFunc("/tareas/delete", handler.DeleteTareaHandler)

	// Aplicar middleware CORS
	corsHandler := CORSMiddleware(mux)

	fmt.Printf("Servidor escuchando en http://*:8080\n")
	http.ListenAndServe(":8080", corsHandler)
}
