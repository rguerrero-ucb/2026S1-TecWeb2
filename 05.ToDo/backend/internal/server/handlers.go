package server

import (
	"encoding/json"
	"fmt"
	"log/slog"
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

// HomeHandler maneja las solicitudes al endpoint raíz y valida rutas no encontradas
func (h *Handler) HomeHandler(w http.ResponseWriter, r *http.Request) {
	// Si la ruta no es exactamente "/", devolver 404
	if r.URL.Path != "/" {
		SendResponse(w, map[string]any{
			"success": false,
			"error":   fmt.Sprintf("Ruta no encontrada: %s", r.URL.Path),
		})
		w.WriteHeader(http.StatusNotFound)
		return
	}

	SendResponse(w, map[string]any{
		"success": true,
		"message": "API de Tareas - Endpoints disponibles: GET /tareas, POST /tareas/create, DELETE /tareas/delete/{id}",
	})
}

// GetAllTareas obtiene todas las tareas
func (h *Handler) GetAllTareasHandler(w http.ResponseWriter, r *http.Request) {
	slog.Info("GetAllTareasHandler llamado", "method", r.Method, "url", r.URL.Path)
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

// CreateTarea crea una nueva tarea
func (h *Handler) CreateTareaHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		SendResponse(w, map[string]any{
			"success": false,
			"error":   "El método debe ser POST",
		})
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

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

// DeleteTarea elimina una tarea
func (h *Handler) DeleteTareaHandler(w http.ResponseWriter, r *http.Request) {
	slog.Info("DeleteTareaHandler llamado", "method", r.Method, "url", r.URL.Path)

	// Extraer el ID de la URL (formato: /tareas/delete/{id})
	parts := strings.Split(strings.TrimPrefix(r.URL.Path, "/tareas/delete/"), "/")
	if len(parts) < 1 || parts[0] == "" {
		SendResponse(w, map[string]any{
			"success": false,
			"error":   "El 'id' es requerido en la URL",
		})
		return
	}

	id, err := strconv.Atoi(parts[0])
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
