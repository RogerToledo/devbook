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
}