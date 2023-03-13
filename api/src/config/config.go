package config

import (
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

var (
	StringConexao = ""
	Porta = 0
)

func Carregar() {

	if erro := godotenv.Load(); erro != nil {
		log.Fatal(erro)
	}

	Porta, _ = strconv.Atoi(os.Getenv("API_PORT"))

	StringConexao = fmt.Sprintf("%s:%s@/%s?charset=utf8&parseTime=True&loc=Local",
		os.Getenv("DB_USUARIO"),
		os.Getenv("DB_SENHA"),
		os.Getenv("DB_NOME"),
	)
}