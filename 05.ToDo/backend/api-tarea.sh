# para probar la API de tareas, se pueden usar los siguientes comandos:

curl http://localhost:8080/tareas

curl http://localhost:8080/tareas/create -d "nombre=tarea 100"

curl http://localhost:8080/tareas/delete -d "id=26"

curl -H "ngrok-skip-browser-warning:1" \
    https://fiercely-ungrimed-alexander.ngrok-free.dev/tareas/create '
    -d "nombre=tarea_nueva"