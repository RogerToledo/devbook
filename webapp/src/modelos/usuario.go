package modelos

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"time"
	"webapp/src/config"
	"webapp/src/requisicoes"
)

type Usuario struct {
	ID          uint64        `json:"id"`
	Nome        string        `json:"nome"`
	Email       string        `json:"email"`
	Nick        string        `json:"nick"`
	CriadoEm    time.Time     `json:"criadoEm"`
	Seguidores  []Usuario     `json:"seguidores"`
	Seguindo    []Usuario     `json:"seguindo"`
	Publicacoes []Publicacoes `json:"publicacoes"`
}

func BuscarUsuarioCompleto(usuarioID uint64, r *http.Request) (Usuario, error) {
	var (
		usuario     Usuario
		seguidores  []Usuario
		seguindo    []Usuario
		publicacoes []Publicacoes
	)

	chUsuario := make(chan Usuario)
	chSeguidores := make(chan []Usuario)
	chSeguindo := make(chan []Usuario)
	chPublicacoes := make(chan []Publicacoes)

	go BuscarDadosUsuario(chUsuario, usuarioID, r)
	go BuscarSeguidores(chSeguidores, usuarioID, r)
	go BuscarSeguindo(chSeguindo, usuarioID, r)
	go BuscarPublicacao(chPublicacoes, usuarioID, r)

	for i := 0; i < 4; i++ {
		select {
		case usuarioCarregado := <-chUsuario:
			if usuarioCarregado.ID == 0 {
				return Usuario{}, errors.New("erro ao buscar usuário")
			}

			usuario = usuarioCarregado

		case seguidoresCarregados := <-chSeguidores:
			if seguidoresCarregados == nil {
				return Usuario{}, errors.New("erro ao buscar seguidores")
			}

			seguidores = seguidoresCarregados

		case seguindoCarregados := <-chSeguindo:
			if seguindoCarregados == nil {
				return Usuario{}, errors.New("erro ao buscar seguindo")
			}

			seguindo = seguindoCarregados

		case publicacoesCarregadas := <-chPublicacoes:
			if publicacoesCarregadas == nil {
				return Usuario{}, errors.New("erro ao buscar publicacoes")
			}

			publicacoes = publicacoesCarregadas
		}
	}

	usuario.Seguidores = seguidores
	usuario.Seguindo = seguindo
	usuario.Publicacoes = publicacoes

	return usuario, nil
}

func BuscarDadosUsuario(canal chan<- Usuario, usuarioID uint64, r *http.Request) {
	url := fmt.Sprintf("%s/usuarios/%d", config.ApiUrl, usuarioID)
	response, erro := requisicoes.FazerRequisicaoAutenticacao(r, http.MethodGet, url, nil)
	if erro != nil {
		canal <- Usuario{}
		return
	}
	defer response.Body.Close()

	var usuario Usuario
	if erro = json.NewDecoder(response.Body).Decode(&usuario); erro != nil {
		canal <- Usuario{}
		return
	}

	canal <- usuario
}

func BuscarSeguidores(canal chan<- []Usuario, usuarioID uint64, r *http.Request) {
	url := fmt.Sprintf("%s/usuarios/listar-seguidores/%d", config.ApiUrl, usuarioID)
	response, erro := requisicoes.FazerRequisicaoAutenticacao(r, http.MethodGet, url, nil)
	if erro != nil {
		canal <- nil
		return
	}
	defer response.Body.Close()

	var seguidores []Usuario
	if erro = json.NewDecoder(response.Body).Decode(&seguidores); erro != nil {
		canal <- nil
		return
	}

	if seguidores == nil {
		canal <- make([]Usuario, 0)
		return
	}

	canal <- seguidores
}

func BuscarSeguindo(canal chan<- []Usuario, usuarioID uint64, r *http.Request) {
	url := fmt.Sprintf("%s/usuarios/listar-seguindo/%d", config.ApiUrl, usuarioID)
	response, erro := requisicoes.FazerRequisicaoAutenticacao(r, http.MethodGet, url, nil)
	if erro != nil {
		canal <- nil
		return
	}
	defer response.Body.Close()

	var seguindo []Usuario
	if erro = json.NewDecoder(response.Body).Decode(&seguindo); erro != nil {
		canal <- nil
		return
	}

	if seguindo == nil {
		canal <- make([]Usuario, 0)
		return
	}

	canal <- seguindo
}

func BuscarPublicacao(canal chan<- []Publicacoes, usuarioID uint64, r *http.Request) {
	url := fmt.Sprintf("%s/publicacoes/usuario/%d", config.ApiUrl, usuarioID)
	response, erro := requisicoes.FazerRequisicaoAutenticacao(r, http.MethodGet, url, nil)
	if erro != nil {
		canal <- nil
		return
	}
	defer response.Body.Close()

	var publicacoes []Publicacoes
	if erro = json.NewDecoder(response.Body).Decode(&publicacoes); erro != nil {
		canal <- nil
		return
	}

	if publicacoes == nil {
		canal <- make([]Publicacoes, 0)
		return
	}

	canal <- publicacoes
}
