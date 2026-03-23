package repository

import (
	"database/sql"
	"log/slog"

	_ "github.com/lib/pq" // PostgreSQL driver
)

func InitDB() *sql.DB {
	// Construye la cadena de conexión
	// Abre la conexión a la base de datos
	db, err := sql.Open("postgres", "host=localhost port=5432 user=ucb password=Tarija00 dbname=ucb sslmode=disable")
	if err != nil {
		slog.Error("Error al abrir la conexión a la base de datos", "error", err)
		return nil
	}

	// Verifica la conexión
	err = db.Ping()
	if err != nil {
		slog.Error("Error al conectar a la base de datos", "error", err)
		return nil
	}

	slog.Info("Conexión a la base de datos establecida exitosamente")
	return db

}
