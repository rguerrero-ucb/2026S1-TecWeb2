package server

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"strings"
	"ucb/todo/internal/service"
)

// Handler es una estructura que contiene las dependencias necesarias para los handlers
type Handler struct {
	tareaService *service.TareaService
}

// NewHandler crea una nueva instancia de Handler
func NewHandler(tareaService *service.TareaService) *Handler {
	return &Handler{
		tareaService: tareaService,
	}
}

// HomeHandler maneja las solicitudes al endpoint raíz
func (h *Handler) HomeHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "¡Hola, mundo!")
}

// TareasHandler maneja GET, POST y DELETE en /tareas
func (h *Handler) TareasHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		h.getAllTareas(w, r)
	case http.MethodPost:
		h.createTarea(w, r)
	case http.MethodDelete:
		h.deleteTarea(w, r)
	default:
		http.Error(w, "Método no permitido", http.StatusMethodNotAllowed)
	}
}

// getAllTareas obtiene todas las tareas
func (h *Handler) getAllTareas(w http.ResponseWriter, r *http.Request) {
	tareas, err := h.tareaService.GetAllTareas()
	if err != nil {
		http.Error(w, "Error al obtener las tareas", http.StatusInternalServerError)
		return
	}

	SendResponse(w, map[string]any{
		"success": true,
		"tareas":  tareas,
	})
}

// createTarea crea una nueva tarea
func (h *Handler) createTarea(w http.ResponseWriter, r *http.Request) {

	var req struct {
		Nombre string `json:"nombre"`
		Estado int    `json:"estado"`
	}

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		SendResponse(w, map[string]any{
			"success": false,
			"error":   "Error al decodificar el JSON",
		})
		return
	}

	if req.Nombre == "" {
		SendResponse(w, map[string]any{
			"success": false,
			"error":   "El campo 'nombre' es requerido",
		})
		return
	}

	tarea, err := h.tareaService.CreateTarea(req.Nombre)
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
}

// deleteTarea elimina una tarea
func (h *Handler) deleteTarea(w http.ResponseWriter, r *http.Request) {
	// Extraer el ID de la URL (formato: /tareas/{id})
	parts := strings.Split(strings.TrimPrefix(r.URL.Path, "/tareas"), "/")
	if len(parts) < 2 || parts[1] == "" {
		SendResponse(w, map[string]any{
			"success": false,
			"error":   "El 'id' es requerido en la URL",
		})
		return
	}

	id, err := strconv.Atoi(parts[1])
	if err != nil {
		SendResponse(w, map[string]any{
			"success": false,
			"error":   "El 'id' debe ser un número válido",
		})
		return
	}

	err = h.tareaService.DeleteTarea(id)
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
}

// GetAllTareasHandler (deprecated) - Mantenido por compatibilidad
func (h *Handler) GetAllTareasHandler(w http.ResponseWriter, r *http.Request) {
	h.getAllTareas(w, r)
}

// CreateTareaHandler (deprecated) - Mantenido por compatibilidad
func (h *Handler) CreateTareaHandler(w http.ResponseWriter, r *http.Request) {
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

	tarea, err := h.tareaService.CreateTarea(nombre)
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
}

// DeleteTareaHandler (deprecated) - Mantenido por compatibilidad
func (h *Handler) DeleteTareaHandler(w http.ResponseWriter, r *http.Request) {
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

	err = h.tareaService.DeleteTarea(id)
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
}
