package rotas

import (
	"api/src/controllers"
	"net/http"
)

var rotasUsuarios = []Rota{
	{
		Uri:                "/usuarios",
		Metodo:             http.MethodPost,
		Funcao:             controllers.CriarUsuario,
		RequerAutenticacao: false,
	},
	{
		Uri:                "/usuarios",
		Metodo:             http.MethodGet,
		Funcao:             controllers.BuscarUsuarios,
		RequerAutenticacao: true,
	},
	{
		Uri:                "/usuarios/{id}",
		Metodo:             http.MethodGet,
		Funcao:             controllers.BuscarUsuario,
		RequerAutenticacao: true,
	},
	{
		Uri:                "/usuarios/{id}",
		Metodo:             http.MethodPut,
		Funcao:             controllers.AtualizarUsuario,
		RequerAutenticacao: true,
	},
	{
		Uri:                "/usuarios/{id}",
		Metodo:             http.MethodDelete,
		Funcao:             controllers.DeletarUsuario,
		RequerAutenticacao: true,
	},
	{
		Uri:                "/usuarios/seguir/{id}",
		Metodo:             http.MethodPost,
		Funcao:             controllers.SeguirUsuario,
		RequerAutenticacao: true,
	},
	{
		Uri:                "/usuarios/parar-seguir/{id}",
		Metodo:             http.MethodPost,
		Funcao:             controllers.PararSeguirUsuario,
		RequerAutenticacao: true,
	},
	{
		Uri:                "/usuarios/listar-seguidores/{id}",
		Metodo:             http.MethodGet,
		Funcao:             controllers.ListarSeguindoSeguidores,
		RequerAutenticacao: true,
	},
	{
		Uri:                "/usuarios/listar-seguindo/{id}",
		Metodo:             http.MethodGet,
		Funcao:             controllers.ListarSeguindoSeguidores,
		RequerAutenticacao: true,
	},
	{
		Uri:                "/usuarios/alterar-senha/{id}",
		Metodo:             http.MethodPost,
		Funcao:             controllers.AlterarSenha,
		RequerAutenticacao: true,
	},
}
