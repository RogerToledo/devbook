package db

import (
	"api/src/config"
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

func Conectar() (*sql.DB, error) {
	conn, erro := sql.Open("mysql", config.StringConexao)
	if erro != nil {
		return nil, erro
	}

	if erro = conn.Ping(); erro != nil {
		conn.Close()
		return nil, erro
	}

	return conn, nil
}