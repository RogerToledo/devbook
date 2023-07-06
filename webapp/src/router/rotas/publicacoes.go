package rotas

import (
	"net/http"
	"webapp/src/controller"
)

var rotasPublicacoes = []Rota{
	{
		URI:                "/publicacoes",
		Metodo:             http.MethodPost,
		Funcao:             controller.Criar,
		RequerAutenticacao: true,
	},
	{
		URI:                "/publicacoes/curtir/{publicacaoID}",
		Metodo:             http.MethodPost,
		Funcao:             controller.Curtir,
		RequerAutenticacao: true,
	},
	{
		URI:                "/publicacoes/editar/{publicacaoID}",
		Metodo:             http.MethodGet,
		Funcao:             controller.CarregaPaginaEdicao,
		RequerAutenticacao: true,
	},
	{
		URI:                "/publicacoes/editar/{publicacaoID}",
		Metodo:             http.MethodPut,
		Funcao:             controller.Atualizar,
		RequerAutenticacao: true,
	},
	{
		URI:                "/publicacoes/{publicacaoID}",
		Metodo:             http.MethodDelete,
		Funcao:             controller.Deletar,
		RequerAutenticacao: true,
	},
}
