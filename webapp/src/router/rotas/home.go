package rotas

import (
	"net/http"
	"webapp/src/controller"
)

var rotaPaginaPrincipal = Rota{
	URI: "/home",
	Metodo: http.MethodGet,
	Funcao: controller.CarregaPaginaPrincipal,
	RequerAutenticacao: true,
}