package controllers

import (
	"api/src/autenticacao"
	"api/src/db"
	"api/src/modelos"
	"api/src/repositorios"
	"api/src/respostas"
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func CriarPublicacoes(w http.ResponseWriter, r *http.Request) {
	corpoRequest, erro := ioutil.ReadAll(r.Body)
	if erro != nil {
		respostas.Erro(w, http.StatusBadRequest, erro)
		return
	}

	var publicacao modelos.Publicacao
	if erro := json.Unmarshal(corpoRequest, &publicacao); erro != nil {
		respostas.Erro(w, http.StatusUnprocessableEntity, erro)
		return
	}

	if erro := publicacao.Preparar(); erro != nil {
		respostas.Erro(w, http.StatusInternalServerError, erro)
		return
	}

	IDToken, erro := autenticacao.ExtrairUsuarioID(r)
	if erro != nil {
		respostas.Erro(w, http.StatusUnauthorized, erro)
	}

	db, erro := db.Conectar()
	if erro != nil {
		respostas.Erro(w, http.StatusInternalServerError, erro)
		return
	}
	defer db.Close()

	publicacao.AutorID = IDToken

	repoPublicacoes := repositorios.NovoRepositorioDePubicacoes(db)
	IDPublicacao, erro := repoPublicacoes.Criar(publicacao)
	if erro != nil {
		respostas.Erro(w, http.StatusInternalServerError, erro)
		return
	}

	publicacao.ID = IDPublicacao

	respostas.Json(w, http.StatusOK, publicacao)
}
// BuscarPublicacoes busca por publicacoes do usuario e dos que ele segue
func BuscarPublicacoes(w http.ResponseWriter, r *http.Request) {
	IDToken, erro := autenticacao.ExtrairUsuarioID(r)
	if erro != nil {
		respostas.Erro(w, http.StatusBadRequest, erro)
		return
	}

	db, erro := db.Conectar()
	if erro != nil {
		respostas.Erro(w, http.StatusInternalServerError, erro)
		return
	}
	defer db.Close()

	repositorio := repositorios.NovoRepositorioDePubicacoes(db)
	publicacoes, erro := repositorio.BuscarPublicacoes(IDToken)
	if erro != nil {
		respostas.Erro(w, http.StatusInternalServerError, erro)
		return
	}

	respostas.Json(w, http.StatusOK, publicacoes)
}

func BuscarPublicacaoID(w http.ResponseWriter, r *http.Request) {
	parametro := mux.Vars(r)

	ID, erro := strconv.ParseUint(parametro["publicacaoId"], 10, 64)
	if erro != nil {
		respostas.Erro(w, http.StatusBadRequest, erro)
		return
	}

	db, erro := db.Conectar()
	if erro != nil {
		respostas.Erro(w, http.StatusInternalServerError, erro)
		return
	}
	defer db.Close()

	repositorio := repositorios.NovoRepositorioDePubicacoes(db)
	publicacao, erro := repositorio.BuscarPorId(ID)
	if erro != nil {
		respostas.Erro(w, http.StatusInternalServerError, erro)
		return
	}

	respostas.Json(w, http.StatusOK, publicacao)	
}

func BuscarPublicacoesUsuario(w http.ResponseWriter, r *http.Request) {
	parametro := mux.Vars(r)
	usuarioID, erro := strconv.ParseUint(parametro["usuarioId"], 10, 64)
	if erro != nil {
		respostas.Erro(w, http.StatusBadRequest, erro)
		return
	}

	db, erro := db.Conectar()
	if erro != nil {
		respostas.Erro(w, http.StatusInternalServerError, erro)
	}
	defer db.Close()

	repositorio := repositorios.NovoRepositorioDePubicacoes(db)
	publicacoes, erro := repositorio.BuscarPorUsuario(usuarioID)
	if erro != nil {
		respostas.Erro(w, http.StatusInternalServerError, erro)
		return
	}

	respostas.Json(w, http.StatusOK, publicacoes)
}

func AtualizarPublicacao(w http.ResponseWriter, r *http.Request) {
	db, erro := db.Conectar()
	if erro != nil {
		respostas.Erro(w, http.StatusInternalServerError, erro)
		return
	}
	defer db.Close()

	tokenID, erro := autenticacao.ExtrairUsuarioID(r)
	if erro != nil {
		respostas.Erro(w, http.StatusBadRequest, erro)
		return
	}

	parametro := mux.Vars(r)
	publicacaoID, erro := strconv.ParseUint(parametro["publicacaoId"], 10, 64)
	if erro != nil {
		respostas.Erro(w, http.StatusBadRequest, erro)
		return
	}

	repositorio := repositorios.NovoRepositorioDePubicacoes(db)

	rp, erro := repositorio.BuscarPorId(publicacaoID)
	if erro != nil {
		respostas.Erro(w, http.StatusInternalServerError, erro)
		return
	}

	if tokenID != rp.AutorID {
		respostas.Erro(w, http.StatusForbidden, errors.New("não é possível alterar publicação de outro usuário"))
		return
	}
	
	var publicacao modelos.Publicacao
	corpoRequest, erro := ioutil.ReadAll(r.Body)
	if erro != nil {
		respostas.Erro(w, http.StatusBadRequest, erro)
		return
	}

	if erro := json.Unmarshal(corpoRequest, &publicacao); erro != nil {
		respostas.Erro(w, http.StatusUnprocessableEntity, erro)
		return
	}

	if erro := publicacao.Preparar(); erro != nil {
		respostas.Erro(w, http.StatusBadRequest, erro)
		return
	}
	
	if erro := repositorio.Alterar(publicacaoID, publicacao); erro != nil {
		respostas.Erro(w, http.StatusInternalServerError, erro)
		return
	}

	respostas.Json(w, http.StatusNoContent, nil)
}

func DeletarPublicacao(w http.ResponseWriter, r *http.Request) {
	tokenID, erro := autenticacao.ExtrairUsuarioID(r)
	if erro != nil {
		respostas.Erro(w, http.StatusBadRequest, erro)
		return
	}

	parametro := mux.Vars(r)
	publicacaoID, erro := strconv.ParseUint(parametro["publicacaoId"], 10, 64)
	if erro != nil {
		respostas.Erro(w, http.StatusBadRequest, erro)
		return
	}

	db, erro := db.Conectar()
	if erro != nil {
		respostas.Erro(w, http.StatusInternalServerError, erro)
		return
	}
	defer db.Close()


	repositorio := repositorios.NovoRepositorioDePubicacoes(db)
	publicacao, erro := repositorio.BuscarPorId(publicacaoID)
	if erro != nil {
		respostas.Erro(w, http.StatusInternalServerError, erro)
		return
	}

	if tokenID != publicacao.AutorID {
		respostas.Erro(w, http.StatusForbidden, errors.New("não é possível deletar publicacoes de outros usuário"))
		return
	}

	if erro := repositorio.Deletar(publicacaoID); erro != nil {
		respostas.Erro(w, http.StatusInternalServerError, erro)
		return
	}
	
	respostas.Json(w, http.StatusNoContent, nil)
}

func CurtirPublicacao(w http.ResponseWriter, r *http.Request) {
	usuarioID, erro := autenticacao.ExtrairUsuarioID(r)
	if erro != nil {
		respostas.Erro(w, http.StatusBadRequest, erro)
		return
	}

	parametro := mux.Vars(r)
	publicacaoID, erro := strconv.ParseUint(parametro["publicacaoId"], 10, 64)
	if erro != nil {
		respostas.Erro(w, http.StatusBadRequest, erro)
		return
	}

	db, erro := db.Conectar()
	if erro != nil {
		respostas.Erro(w, http.StatusInternalServerError, erro)
		return
	}
	defer db.Close()

	repositorio := repositorios.NovoRepositorioDePubicacoes(db)

	publicacaoSalva, erro := repositorio.BuscarPorId(publicacaoID)
	if erro != nil {
		respostas.Erro(w, http.StatusInternalServerError, erro)
		return
	}

	if publicacaoSalva.ID == 0 {
		respostas.Erro(w, http.StatusBadRequest, errors.New("publicação não existe"))
		return
	}

	if erro := repositorio.Curtir(publicacaoID, usuarioID); erro != nil {
		respostas.Erro(w, http.StatusInternalServerError, erro)
		return
	}

	respostas.Json(w, http.StatusNoContent, nil)
}

func DeixarCurtirPublicacao(w http.ResponseWriter, r *http.Request) {
	tokenID, erro := autenticacao.ExtrairUsuarioID(r)
	if erro != nil {
		respostas.Erro(w, http.StatusBadRequest, erro)
		return
	}

	parametro := mux.Vars(r)
	publicacaoID, erro := strconv.ParseUint(parametro["publicacaoId"], 10, 64)
	if erro != nil {
		respostas.Erro(w, http.StatusBadRequest, erro)
		return
	}

	db, erro := db.Conectar()
	if erro != nil {
		respostas.Erro(w, http.StatusInternalServerError, erro)
		return
	}
	defer db.Close()

	repositorio := repositorios.NovoRepositorioDePubicacoes(db)
	if erro := repositorio.DeixarCurtir(publicacaoID, tokenID); erro != nil {
		respostas.Erro(w, http.StatusInternalServerError, erro)
		return
	}

	respostas.Json(w, http.StatusNoContent, nil)
}
