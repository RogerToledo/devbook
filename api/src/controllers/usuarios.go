package controllers

import (
	"api/src/autenticacao"
	"api/src/db"
	"api/src/modelos"
	"api/src/repositorios"
	"api/src/respostas"
	"api/src/seguranca"
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"

	"github.com/gorilla/mux"
)

func CriarUsuario(w http.ResponseWriter, r *http.Request) {
	corpoRequest, erro := ioutil.ReadAll(r.Body)
	if erro != nil {
		respostas.Erro(w, http.StatusUnprocessableEntity, erro)
		return
	}

	var usuario modelos.Usuario
	if erro = json.Unmarshal(corpoRequest, &usuario); erro != nil {
		respostas.Erro(w, http.StatusBadRequest, erro)
		return
	}

	if erro := usuario.Prepare("cadastro"); erro != nil {
		respostas.Erro(w, http.StatusBadRequest, erro)
		return
	}

	db, erro := db.Conectar()
	if erro != nil {
		respostas.Erro(w, http.StatusInternalServerError, erro)
		return
	}
	defer db.Close()

	repositorio := repositorios.NovoRepositorioDeUsuarios(db)
	usuario.ID, erro = repositorio.Criar(usuario)
	if erro != nil {
		respostas.Erro(w, http.StatusInternalServerError, erro)
		return
	}

	usuarioRet := struct {
		ID   int32  `json:"id"`
		Nome string `json:"nome"`
	}{
		ID:   int32(usuario.ID),
		Nome: usuario.Nome,
	}

	respostas.Json(w, http.StatusOK, usuarioRet)
}

func BuscarUsuarios(w http.ResponseWriter, r *http.Request) {
	usuario := strings.ToLower(r.URL.Query().Get("usuario"))

	db, erro := db.Conectar()
	if erro != nil {
		respostas.Erro(w, http.StatusInternalServerError, erro)
	}
	defer db.Close()

	repositorio := repositorios.NovoRepositorioDeUsuarios(db)

	if usuario != "" {
		usuarios, erro := repositorio.BuscarPorNomeNick(usuario)
		if erro != nil {
			respostas.Erro(w, http.StatusInternalServerError, erro)
			return
		}
	
		respostas.Json(w, http.StatusOK, usuarios)

	} else {
		usuarios, erro := repositorio.Buscar()
		if erro != nil {
			respostas.Erro(w, http.StatusInternalServerError, erro)
			return
		}
	
		respostas.Json(w, http.StatusOK, usuarios)
	}
	
	

}

func BuscarUsuario(w http.ResponseWriter, r *http.Request) {
	parametro := mux.Vars(r)

	ID, erro := strconv.ParseUint(parametro["id"], 10, 64)
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

	repositorio := repositorios.NovoRepositorioDeUsuarios(db)
	usuario, erro := repositorio.BuscarPorId(ID)
	if erro != nil {
		respostas.Erro(w, http.StatusInternalServerError, erro)
		return
	}

	if usuario.ID == 0 && usuario.Nome == "" {
		respostas.Erro(w, http.StatusOK, errors.New("usuario não encontrado"))
		return
	}

	respostas.Json(w, http.StatusOK, usuario)
}

func AtualizarUsuario(w http.ResponseWriter, r *http.Request) {
	parametro := mux.Vars(r)

	ID, erro := strconv.ParseUint(parametro["id"], 10, 64)
	if erro != nil {
		respostas.Erro(w, http.StatusBadRequest, erro)
		return
	}

	IDToken, erro := autenticacao.ExtrairUsuarioID(r)
	if erro != nil {
		respostas.Erro(w, http.StatusUnauthorized, erro)
		return
	}

	if ID != IDToken {
		respostas.Erro(w, http.StatusForbidden, errors.New("não é possível atualizar outro usuário"))
		return
	}

	corpoRequisicao, erro := ioutil.ReadAll(r.Body)
	if erro != nil {
		respostas.Erro(w, http.StatusBadRequest, erro)
		return
	}

	var usuario modelos.Usuario
	if erro := json.Unmarshal(corpoRequisicao, &usuario); erro != nil {
		respostas.Erro(w, http.StatusBadRequest, erro)
		return
	}

	if erro := usuario.Prepare("atualizar"); erro != nil {
		respostas.Erro(w, http.StatusInternalServerError, erro)
		return
	}

	db, erro := db.Conectar()
	if erro != nil {
		respostas.Erro(w, http.StatusInternalServerError, erro)
		return
	}
	defer db.Close()

	repositorio := repositorios.NovoRepositorioDeUsuarios(db)
	if erro := repositorio.Atualizar(ID, usuario); erro != nil {
		respostas.Erro(w, http.StatusInternalServerError, erro)
		return
	}

	respostas.Json(w, http.StatusNoContent, nil)
}

func DeletarUsuario(w http.ResponseWriter, r *http.Request) {
	parametro := mux.Vars(r)

	ID, erro := strconv.ParseUint(parametro["id"], 10, 64)
	if erro != nil {
		respostas.Erro(w, http.StatusBadRequest, erro)
	}

	IDToken, erro := autenticacao.ExtrairUsuarioID(r)
	if erro != nil {
		respostas.Erro(w, http.StatusUnauthorized, erro)
		return
	}

	if ID != IDToken {
		respostas.Erro(w, http.StatusForbidden, errors.New("não é possível deletar outro usuário"))
		return
	}

	db, erro := db.Conectar()
	if erro != nil {
		respostas.Erro(w, http.StatusInternalServerError, erro)
	}
	defer db.Close()

	repositorio := repositorios.NovoRepositorioDeUsuarios(db)
	if erro := repositorio.Deletar(ID); erro != nil {
		respostas.Erro(w, http.StatusInternalServerError, erro)
	}
}

func SeguirUsuario(w http.ResponseWriter, r *http.Request) {
	parametro := mux.Vars(r)

	ID, erro := strconv.ParseUint(parametro["id"], 10, 64)
	if erro != nil {
		respostas.Erro(w, http.StatusBadRequest, erro)
		return
	}

	seguidorID, erro := autenticacao.ExtrairUsuarioID(r)
	if erro != nil {
		respostas.Erro(w, http.StatusForbidden, erro)
		return
	}

	if seguidorID == ID {
		respostas.Erro(w, http.StatusForbidden, errors.New("não é possível seguir a si mesmo"))
		return
	}

	db, erro := db.Conectar()
	if erro != nil {
		respostas.Erro(w, http.StatusInternalServerError, erro)
		return
	}
	defer db.Close()

	repositorio := repositorios.NovoRepositorioDeUsuarios(db)

	if erro := repositorio.Seguir(seguidorID, ID); erro != nil {
		respostas.Erro(w, http.StatusInternalServerError, erro)
		return
	}

	respostas.Json(w, http.StatusNoContent, nil)
}

func PararSeguirUsuario(w http.ResponseWriter, r *http.Request) {
	parametros := mux.Vars(r)

	ID, erro := strconv.ParseUint(parametros["id"], 10, 64)
	if erro != nil {
		respostas.Erro(w, http.StatusBadRequest, erro)
	}

	seguidorID, erro := autenticacao.ExtrairUsuarioID(r)
	if erro != nil {
		respostas.Erro(w, http.StatusForbidden, erro)
	}

	db, erro := db.Conectar()
	if erro != nil {
		respostas.Erro(w, http.StatusInternalServerError, erro)
	}
	defer db.Close()

	repositorio := repositorios.NovoRepositorioDeUsuarios(db)
	if erro := repositorio.PararSeguirUsuario(seguidorID, ID); erro != nil {
		respostas.Erro(w, http.StatusInternalServerError, erro)
	}

	respostas.Json(w, http.StatusNoContent, nil)
}

func ListarSeguindoSeguidores(w http.ResponseWriter, r *http.Request) {
	uri := strings.Split(r.RequestURI, "/")

	tipo := uri[2]

	parametros := mux.Vars(r)

	ID, erro := strconv.ParseUint(parametros["id"], 10, 64)
	if erro != nil {
		respostas.Erro(w, http.StatusBadRequest, erro)
	}

	db, erro := db.Conectar()
	if erro != nil {
		respostas.Erro(w, http.StatusInternalServerError, erro)
	}
	defer db.Close()

	repositorio := repositorios.NovoRepositorioDeUsuarios(db)

	seguidores, erro := repositorio.ListarSeguindoSeguidores(ID, tipo)
	if erro != nil {
		respostas.Erro(w, http.StatusInternalServerError, erro)
	}

	respostas.Json(w, http.StatusOK, seguidores)
}

func AlterarSenha(w http.ResponseWriter, r *http.Request) {
	corpoRequest, erro := ioutil.ReadAll(r.Body)
	if erro != nil {
		respostas.Erro(w, http.StatusUnprocessableEntity, erro)
		return
	}

	var senhas modelos.AlterarSenha
	if erro := json.Unmarshal(corpoRequest, &senhas); erro != nil {
		respostas.Erro(w, http.StatusInternalServerError, erro)
		return
	}

	if erro := senhas.ValidarSenhas(); erro != nil {
		respostas.Erro(w, http.StatusBadRequest, erro)
		return
	}

	parametros := mux.Vars(r)

	ID, erro := strconv.ParseUint(parametros["id"], 10, 64)
	if erro != nil {
		respostas.Erro(w, http.StatusBadRequest, erro)
		return
	}

	IDToken, erro := autenticacao.ExtrairUsuarioID(r)
	if erro != nil {
		respostas.Erro(w, http.StatusInternalServerError, erro)
		return
	}

	if ID != IDToken {
		respostas.Erro(w, http.StatusForbidden, errors.New("só é permitido alterar sua própria senha"))
		return
	}

	db, erro := db.Conectar()
	if erro != nil {
		respostas.Erro(w, http.StatusInternalServerError, erro)
		return
	}

	repositorio := repositorios.NovoRepositorioDeUsuarios(db)
	senhaSalva, erro := repositorio.BuscarSenha(ID)
	if erro != nil {
		respostas.Erro(w, http.StatusInternalServerError, erro)
		return
	}

	if erro := seguranca.VerificaSenha(senhaSalva, senhas.Atual); erro != nil {
		respostas.Erro(w, http.StatusUnauthorized, errors.New("senha atual incorreta"))
		return
	}

	senhaHash, erro := seguranca.Hash(senhas.Nova)
	if erro != nil {
		respostas.Erro(w, http.StatusBadRequest, erro)
		return
	}

	if erro := repositorio.AlterarSenha(ID, string(senhaHash)); erro != nil {
		respostas.Erro(w, http.StatusInternalServerError, erro)
		return
	}

	respostas.Json(w, http.StatusNoContent, nil)
}
