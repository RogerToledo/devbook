package modelos

import (
	"testing"
)

func TestValidar(t *testing.T) {
	var cenarios = []struct {
		descricao string
		entregue Usuario
		esperado string
	}{
		{
			descricao: "Deve retornar erro todos os campos estão vazio",
			entregue: Usuario{
				Nome: "",
				Nick: "",
				Email: "",
			},
			esperado: "o nome é obrigatório",
		},{
			descricao: "Deve retornar erro quando Nome está vazio",
			entregue: Usuario{
				Nome: "",
				Nick: "Teste",
				Email: "teste@mail.com",
			},
			esperado: "o nome é obrigatório",
		},{
			descricao: "Deve retornar erro quando Nick está vazio",
			entregue: Usuario{
				Nome: "Teste",
				Nick: "",
				Email: "teste@mail.com",
			},
			esperado: "o nick é obrigatório",
		},{
			descricao: "Deve retornar erro quando Email está vazio",
			entregue: Usuario{
				Nome: "Teste",
				Nick: "Teste",
				Email: "",
			},
			esperado: "o e-mail é obrigatório",
		},
	}

	for _, cenario := range cenarios {
		t.Log(cenario.descricao)

		if erro := cenario.entregue.validar("cadastro"); erro.Error() != cenario.esperado {
			t.Fatalf("Esperado %v, Entregue %v", cenario.esperado, erro)
		}
	}
}

func TestFormatar(t *testing.T) {
	usuarioEsperado := Usuario{
		Nome: "Teste",
		Nick: "Teste",
		Email: "teste@mail.com",
	}

	var cenarios = []struct {
		descricao string
		entregue Usuario
	}{
		{
			descricao: "Deve retornar erro quando Nome não é formatado",
			entregue: Usuario{
				Nome: "  Teste  ",
				Nick: "Teste",
				Email: "teste@mail.com",
			},
		},
		{
			descricao: "Deve retornar erro quando Nick não é formatado",
			entregue: Usuario{
				Nome: "Teste",
				Nick: " Teste   ",
				Email: "teste@mail.com",
			},
		},
		{
			descricao: "Deve retornar erro quando Email não é formatado",
			entregue: Usuario{
				Nome: "Teste",
				Nick: " Teste",
				Email: "teste@mail.com      ",
			},
		},
	}

	for _, cenario := range cenarios {
		t.Log(cenario.descricao)

		if cenario.entregue.formatar("cadastro"); 
			cenario.entregue.Nome != usuarioEsperado.Nome &&
			cenario.entregue.Nick == usuarioEsperado.Nick &&
			cenario.entregue.Email == usuarioEsperado.Email {
				t.Fatalf("Esperado %v, Entregue %v", usuarioEsperado.Nome, cenario.entregue.Nome)
		}

		if cenario.entregue.formatar("cadastro"); 
			cenario.entregue.Nome == usuarioEsperado.Nome &&
			cenario.entregue.Nick != usuarioEsperado.Nick &&
			cenario.entregue.Email == usuarioEsperado.Email {
				t.Fatalf("Esperado %v, Entregue %v", usuarioEsperado.Nick, cenario.entregue.Nick)
		}

		if cenario.entregue.formatar("cadastro"); 
			cenario.entregue.Nome == usuarioEsperado.Nome &&
			cenario.entregue.Nick == usuarioEsperado.Nick &&
			cenario.entregue.Email != usuarioEsperado.Email {
				t.Fatalf("Esperado %v, Entregue %v", usuarioEsperado.Email, cenario.entregue.Email)
		}
	}
	
}