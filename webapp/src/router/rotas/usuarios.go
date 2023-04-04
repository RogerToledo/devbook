package rotas

import (
	"net/http"
	"webapp/src/controller"
)

var rotasUsuarios = []Rota {
	{
		URI: "/criar-usuario",
		Metodo: http.MethodGet,
		Funcao: controller.CarregaCadastroUsuario,
		RequerAutenticacao: false,
	},
	{
		URI: "/usuarios",
		Metodo: http.MethodPost,
		Funcao: controller.CriarUsuario,
		RequerAutenticacao: false,
	},
}