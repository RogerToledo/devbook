package main

import (
	"fmt"
	"log"
	"net/http"
	"webapp/src/router"
	"webapp/src/utils"
)

func main() {
	fmt.Println("Escutando porta 3000")
	r := router.Gerar()
	
	utils.CarregarTemplates()
	log.Fatal(http.ListenAndServe(":3000", r))
}
