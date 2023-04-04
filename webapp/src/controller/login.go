package controller

import (
	"net/http"
	"webapp/src/utils"
)

func CarregarTelaLogin(w http.ResponseWriter, r *http.Request) {
	utils.ExecutarTemplate(w, "login.html", nil)
}