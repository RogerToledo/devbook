package rotas

import (
	"net/http"
	"webapp/src/controller"
)

var rotasPublicacoes = []Rota{
	{
		URI:                "/publicacoes",
		Metodo:             http.MethodPost,
		Funcao:             controller.CriarPublicacao,
		RequerAutenticacao: true,
	},
	{
		URI: "/publicacoes/curtir/{publicacaoID}",
		Metodo: http.MethodPost,
		Funcao: controller.CurtirPublicacao,
		RequerAutenticacao: true,
	},
}
