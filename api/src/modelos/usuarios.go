package modelos

import (
	"errors"
	"strings"
	"time"
	"api/src/seguranca"

	"github.com/badoux/checkmail"
)

type Usuario struct {
	ID       uint64    `json:"id,omitempty"`
	Nome     string    `json:"nome,omitempty"`
	Nick     string    `json:"nick,omitempty"`
	Email    string    `json:"email,omitempty"`
	Senha    string    `json:"senha,omitempty"`
	CriadoEm time.Time `json:"CriadoEm,omitempty"`
}

func (usuario *Usuario) Prepare(etapa string) error {
	if erro := usuario.validar(etapa); erro != nil {
		return erro
	}

	usuario.formatar(etapa)

	return nil
}

func (usuario *Usuario) validar(etapa string) error {
	if usuario.Nome == "" {
		return errors.New("o nome é obrigatório")
	}

	if usuario.Nick == "" {
		return errors.New("o nick é obrigatório")
	}

	if usuario.Email == "" {
		return errors.New("o e-mail é obrigatório")
	}
	if erro := checkmail.ValidateFormat(usuario.Email); erro != nil {
		return errors.New("formato do e-mail é invalido")
	}

	if usuario.Senha == "" && etapa == "cadastro"{
		return errors.New("o senha é obrigatório")
	}

	return nil
}

func (usuario *Usuario) formatar(etapa string) error {
	usuario.Nome  = strings.TrimSpace(usuario.Nome)
	usuario.Nick  = strings.TrimSpace(usuario.Nick)
	usuario.Email = strings.TrimSpace(usuario.Email)

	if etapa == "cadastro" {
		senhaHash, erro := seguranca.Hash(usuario.Senha)
		if erro != nil {
			return erro
		}

		usuario.Senha = string(senhaHash)
	}

	return nil
}
