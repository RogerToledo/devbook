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
	{
		URI: "/buscar-usuarios",
		Metodo: http.MethodGet,
		Funcao: controller.CarregaPaginaUsuarios,
		RequerAutenticacao: false,
	},
	{
		URI: "/usuarios/{usuarioId}",
		Metodo: http.MethodGet,
		Funcao: controller.CarregaPerfilUsuario,
		RequerAutenticacao: false,
	},
	{
		URI: "/usuarios/parar-seguir/{usuarioId}",
		Metodo: http.MethodPost,
		Funcao: controller.PararSeguirUsuario,
		RequerAutenticacao: true,
	},
	{
		URI: "/usuarios/seguir/{usuarioId}",
		Metodo: http.MethodPost,
		Funcao: controller.SeguirUsuario,
		RequerAutenticacao: true,
	},
	{
		URI: "/perfil",
		Metodo: http.MethodGet,
		Funcao: controller.CarregaPerfilUsuarioLogado,
		RequerAutenticacao: true,
	},
	{
		URI: "/editar-usuario",
		Metodo: http.MethodGet,
		Funcao: controller.CarregaPaginaEdicaoUsuario,
		RequerAutenticacao: true,
	},
	{
		URI: "/editar-usuario",
		Metodo: http.MethodPut,
		Funcao: controller.EditarUsuario,
		RequerAutenticacao: true,
	},
	{
		URI: "/atualizar-senha",
		Metodo: http.MethodGet,
		Funcao: controller.CarregarPaginaAtualizarSenha,
		RequerAutenticacao: true,
	},
	{
		URI: "/atualizar-senha",
		Metodo: http.MethodPost,
		Funcao: controller.AtualizarSenha,
		RequerAutenticacao: true,
	},
	{
		URI: "/deletar-usuario",
		Metodo: http.MethodDelete,
		Funcao: controller.DeletarUsuario,
		RequerAutenticacao: true,
	},
}