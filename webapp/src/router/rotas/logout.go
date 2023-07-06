package rotas

import (
	"net/http"
	"webapp/src/controller"
)

var rotasLogout = Rota{
	URI:                "/logout",
	Metodo:             http.MethodGet,
	Funcao:             controller.Logout,
	RequerAutenticacao: true,
}
