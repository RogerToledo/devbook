package modelos

import "errors"

type AlterarSenha struct {
	Atual string `json:"atual"`
	Nova  string `json:"nova"`
}

func (Senhas *AlterarSenha) ValidarSenhas() error {
	if Senhas.Atual == "" {
		return errors.New("a senha atual é obrigatório")
	}

	if Senhas.Nova == "" {
		return errors.New("a senha nova não pode ser vazia")
	}

	return nil
}