package modelos

import (
	"errors"
	"strings"
	"time"
)

type Usuario struct {
	Id       uint64    `json:"id,omitempty"`
	Nome     string    `json:"nome,omitempty"`
	Nick     string    `json:"nick,omitempty"`
	Email    string    `json:"email,omitempty"`
	Senha    string    `json:"senha,omitempty"`
	CriadoEm time.Time `json:"CriadoEm,omitempty"`
}

func (usuario *Usuario) Prepare() error {
	if erro := usuario.validar(); erro != nil {
		return erro
	}

	usuario.formatar()

	return nil
}

func (usuario *Usuario) validar() error {
	if usuario.Nome == "" {
		return errors.New("o nome é obrigatório!!")
	}

	if usuario.Nick == "" {
		return errors.New("o nick é obrigatório!!")
	}

	if usuario.Email == "" {
		return errors.New("o e-mail é obrigatório!!")
	}

	if usuario.Nome == "" {
		return errors.New("o senha é obrigatório!!")
	}

	return nil
}

func (usuario *Usuario) formatar() {
	usuario.Nome  = strings.TrimSpace(usuario.Nome)
	usuario.Nick  = strings.TrimSpace(usuario.Nick)
	usuario.Email = strings.TrimSpace(usuario.Email)
}
