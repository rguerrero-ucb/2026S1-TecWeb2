# Contexto del Proyecto ToDo

## Descripción General
Proyecto de aplicación de gestión de tareas (ToDo) con frontend web simple y backend en Go.

## Estructura del Proyecto
```
/home/ronal/Documentos/UCB/Materias/2026S1.TecnologiasWebII/Temas/05.ToDo/
├── frontend/          # Frontend web (HTML/JavaScript)
│   └── index.html    # Aplicación web principal
└── backend/          # Backend en Go
    └── internal/service/TareaService.go  # Servicio de tareas
```

## Frontend (HTML/JavaScript)
**Archivo:** `frontend/index.html`

### Funcionalidades
1. **Cargar tareas**: Obtiene la lista de tareas desde la API al cargar la página
2. **Agregar tareas**: Permite crear nuevas tareas con nombre
3. **Eliminar tareas**: Elimina tareas existentes con confirmación

### API Endpoints
- **URL base:** `https://fiercely-ungrimed-alexander.ngrok-free.dev/tareas`
- **GET** `/tareas` - Obtener todas las tareas
- **POST** `/tareas` - Crear nueva tarea
- **DELETE** `/tareas/{id}` - Eliminar tarea por ID

### Estructura del Código
- HTML básico con JavaScript integrado
- Usa Fetch API para comunicarse con el backend
- Manejo de errores básico con alertas
- Interfaz minimalista con lista de tareas

## Backend (Go)
**Archivo referenciado:** `backend/internal/service/TareaService.go`

### Contexto del Backend
Según la selección del usuario, el backend utiliza un `tareaRepository` en el servicio de tareas. Esto sugiere una arquitectura por capas:
- **Repository**: Manejo de datos/persistencia
- **Service**: Lógica de negocio
- **Controller/Handler**: Manejo de HTTP

### Tecnologías Inferidas
- **Go** con estructura de proyecto estándar (`internal/`)
- **Repository pattern** para acceso a datos
- **API REST** para comunicación con frontend

## Dependencias y Configuración
- **ngrok**: Usado para exponer el backend localmente
- **Headers especiales**: `ngrok-skip-browser-warning: '1'` para evitar advertencias
- **CORS**: Implícitamente configurado para permitir peticiones del frontend

## Estado Actual
El frontend está funcional y se conecta a un backend expuesto via ngrok. El backend parece seguir buenas prácticas de arquitectura en Go con separación de responsabilidades.

## Notas de Desarrollo
- El frontend es minimalista y podría mejorarse con CSS
- No hay manejo de estados de tareas (completado/pendiente) en el frontend actual
- La comunicación API incluye manejo básico de errores
- El backend sigue patrones comunes de Go (repository, service)

---
*Última actualización: 2026-04-01*
*Contexto registrado durante sesión de desarrollo*