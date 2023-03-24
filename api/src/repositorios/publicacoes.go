package repositorios

import (
	"api/src/modelos"
	"database/sql"
)

type publicacoes struct {
	db *sql.DB
}

func NovoRepositorioDePubicacoes(db *sql.DB) *publicacoes{
	return &publicacoes{db}
}

func (repositorio publicacoes) Criar(publicacao modelos.Publicacao) (uint64, error) {
	ps, erro := repositorio.db.Prepare(
		"insert into publicacoes (titulo, conteudo, autor_id, autor_nick) values (?,?,?,?)")
	if erro != nil {
		return 0, erro
	}
	defer ps.Close()

	resultado, erro := ps.Exec(
		publicacao.Titulo, 
		publicacao.Conteudo, 
		publicacao.AutorID, 
		publicacao.AutorNick,
	)
	if erro != nil {
		return 0, erro
	}

	inserido, erro := resultado.LastInsertId()
	if erro != nil {
		return 0, erro
	}

	return uint64(inserido), nil
}