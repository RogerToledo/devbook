package main

import (
	"fmt"
	"log"

	"net/http"
	"webapp/src/config"
	"webapp/src/cookies"
	"webapp/src/router"
	"webapp/src/utils"
)

func main() {
	r := router.Gerar()

	config.Carregar()
	cookies.ConfigurarCookie()
	utils.CarregarTemplates()
	fmt.Printf("Escutando porta %d", config.Porta)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", config.Porta), r))
}
