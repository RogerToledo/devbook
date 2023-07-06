package controller

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"webapp/src/config"
	"webapp/src/cookies"
	"webapp/src/requisicoes"
	"webapp/src/respostas"

	"github.com/gorilla/mux"
)

func CriarUsuario(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()

	usuario, erro := json.Marshal(map[string]string{
		"nome":  r.FormValue("nome"),
		"nick":  r.FormValue("nick"),
		"email": r.FormValue("email"),
		"senha": r.FormValue("senha"),
	})
	if erro != nil {
		respostas.JSON(w, http.StatusBadRequest, respostas.ErroAPI{Erro: erro.Error()})
		return
	}

	url := fmt.Sprintf("%s/usuarios", config.ApiUrl)
	response, erro := http.Post(url, "application/json", bytes.NewBuffer(usuario))
	if erro != nil {
		respostas.JSON(w, http.StatusInternalServerError, respostas.ErroAPI{Erro: erro.Error()})
		return
	}
	defer response.Body.Close()

	if response.StatusCode >= 400 {
		respostas.TratarStatusCodeErro(w, r, response)
		return
	}

	respostas.JSON(w, response.StatusCode, nil)
}

func SeguirUsuario(w http.ResponseWriter, r *http.Request) {
	parametros := mux.Vars(r)
	usuarioID, erro := strconv.ParseUint(parametros["usuarioId"], 10, 64)
	if erro != nil {
		respostas.JSON(w, http.StatusInternalServerError, respostas.ErroAPI{Erro: erro.Error()})
		return
	}

	url := fmt.Sprintf("%s/usuarios/seguir/%d", config.ApiUrl, usuarioID)
	response, erro := requisicoes.FazerRequisicaoAutenticacao(r, http.MethodPost, url, nil)
	if erro != nil {
		respostas.JSON(w, http.StatusBadRequest, respostas.ErroAPI{Erro: erro.Error()})
		return
	}
	defer r.Body.Close()

	if response.StatusCode >= 400 {
		respostas.TratarStatusCodeErro(w, r, response)
		return
	}

	respostas.JSON(w, response.StatusCode, nil)
}

func PararSeguirUsuario(w http.ResponseWriter, r *http.Request) {
	parametros := mux.Vars(r)
	usuarioID, erro := strconv.ParseUint(parametros["usuarioId"], 10, 64)
	if erro != nil {
		respostas.JSON(w, http.StatusInternalServerError, respostas.ErroAPI{Erro: erro.Error()})
		return
	}

	url := fmt.Sprintf("%s/usuarios/parar-seguir/%d", config.ApiUrl, usuarioID)
	response, erro := requisicoes.FazerRequisicaoAutenticacao(r, http.MethodPost, url, nil)
	if erro != nil {
		respostas.JSON(w, http.StatusBadRequest, respostas.ErroAPI{Erro: erro.Error()})
		return
	}
	defer r.Body.Close()

	if response.StatusCode >= 400 {
		respostas.TratarStatusCodeErro(w, r, response)
		return
	}

	respostas.JSON(w, response.StatusCode, nil)
}

func EditarUsuario(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	usuario, erro := json.Marshal(map[string]string{
		"nome": r.FormValue("nome"),
		"nick": r.FormValue("nick"),
		"email": r.FormValue("email"),
	})
	if erro != nil {
		respostas.JSON(w, http.StatusBadRequest, respostas.ErroAPI{Erro: erro.Error()})
		return
	}

	cookie, _ := cookies.Ler(r)
	usuarioID, _ := strconv.ParseUint(cookie["id"], 10, 64)

	url := fmt.Sprintf("%s/usuarios/%d", config.ApiUrl, usuarioID)

	response, erro := requisicoes.FazerRequisicaoAutenticacao(r, http.MethodPut, url, bytes.NewBuffer(usuario))
	if erro != nil {
		respostas.JSON(w, http.StatusInternalServerError, respostas.ErroAPI{Erro: erro.Error()})
		return
	}
	defer response.Body.Close()

	if response.StatusCode >= 400 {
		respostas.TratarStatusCodeErro(w, r, response)
		return
	}

	respostas.JSON(w, response.StatusCode, nil)
}

func AtualizarSenha(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	senhas, erro := json.Marshal(map[string]string{
		"atual": r.FormValue("senhaAtual"),
		"nova": r.FormValue("novaSenha"),
	})
	if erro != nil {
		respostas.JSON(w, http.StatusBadRequest, respostas.ErroAPI{Erro: erro.Error()})
		return
	}

	cookie, _ := cookies.Ler(r)
	usuarioID, _ := strconv.ParseUint(cookie["id"], 10, 64)
	
	url := fmt.Sprintf("%s/usuarios/alterar-senha/%d", config.ApiUrl, usuarioID)

	response, erro := requisicoes.FazerRequisicaoAutenticacao(r, http.MethodPost, url, bytes.NewBuffer(senhas))
	if erro != nil {
		respostas.JSON(w, http.StatusInternalServerError, respostas.ErroAPI{Erro: erro.Error()})
		return 
	}
	defer response.Body.Close()

	respostas.JSON(w, response.StatusCode, nil)
}	
