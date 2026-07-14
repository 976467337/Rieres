package main

import (
	"database/sql"
	"log"
	"os"
	"path/filepath"

	_ "github.com/jackc/pgx/v5/stdlib"
	"github.com/pressly/goose/v3"
	"github.com/smartkiki/api/internal/config"
)

func main() {
	if len(os.Args) < 2 {
		log.Fatalf("uso: migrate <comando> [args]")
	}

	command := os.Args[1]
	args := os.Args[2:]

	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("falha ao carregar configuracao: %v", err)
	}

	db, err := sql.Open("pgx", cfg.Database.GetDSN())
	if err != nil {
		log.Fatalf("falha ao abrir conexao: %v", err)
	}
	defer db.Close()

	if err := db.Ping(); err != nil {
		log.Fatalf("falha ao conectar ao banco: %v", err)
	}

	if err := goose.SetDialect("postgres"); err != nil {
		log.Fatalf("falha ao configurar dialecto: %v", err)
	}

	migrationsDir, err := filepath.Abs("migrations")
	if err != nil {
		log.Fatalf("falha ao resolver caminho das migracoes: %v", err)
	}

	if err := goose.Run(command, db, migrationsDir, args...); err != nil {
		log.Fatalf("goose %s falhou: %v", command, err)
	}

	log.Printf("comando goose '%s' executado com sucesso", command)
}
