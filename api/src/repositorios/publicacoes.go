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
		"insert into publicacoes (titulo, conteudo, autor_id) values (?,?,?)")
	if erro != nil {
		return 0, erro
	}
	defer ps.Close()

	resultado, erro := ps.Exec(
		publicacao.Titulo, 
		publicacao.Conteudo, 
		publicacao.AutorID,
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

func (repositorio *publicacoes) BuscarPorId(ID uint64) (modelos.Publicacao, error) {
	linha, erro := repositorio.db.Query(
		`select p.id, p.titulo, p.conteudo, p.autor_id, p.curtidas, p.criadaEm, p.autor_id, u.nick 
		from publicacoes p 
		inner join usuarios u on p.autor_id = u.id 
		where p.id = ?`,ID)
	if erro != nil {
		return modelos.Publicacao{}, erro
	}
	defer linha.Close()

	var publicacao modelos.Publicacao
	if linha.Next() {
		if erro := linha.Scan(
			&publicacao.ID,
			&publicacao.Titulo,
			&publicacao.Conteudo,
			&publicacao.AutorID,
			&publicacao.Curtidas,
			&publicacao.CriadaEm,
			&publicacao.AutorID,
			&publicacao.AutorNick,
		); erro != nil {
			return modelos.Publicacao{}, erro
		}
	}
	return publicacao, nil
}

func (repositorio *publicacoes) BuscarPublicacoes(ID uint64) ([]modelos.Publicacao, error) {
	linhas, erro := repositorio.db.Query(
		`select distinct p.id, p.titulo, p.conteudo, p.curtidas, p.criadaEm, p.autor_id, u.nick
		from publicacoes p
		inner join usuarios u 
			on p.autor_id = u.id 
		inner join seguidores s 
			on p.autor_id = s.usuario_id 
		where u.id = ? or s.seguidor_id = ?
		order by p.titulo`, ID, ID,
	)
	if erro != nil {
		return []modelos.Publicacao{}, erro
	}
	defer linhas.Close()

	var publicacoes []modelos.Publicacao
	for linhas.Next() {
		var publicacao modelos.Publicacao
		if erro := linhas.Scan(
			&publicacao.ID,
			&publicacao.Titulo,
			&publicacao.Conteudo,
			&publicacao.Curtidas,
			&publicacao.CriadaEm,
			&publicacao.AutorID,
			&publicacao.AutorNick,
		);erro != nil {
			return []modelos.Publicacao{}, erro
		}

		publicacoes = append(publicacoes, publicacao)
	}
	
	return publicacoes, nil
}

func (repositorio *publicacoes) AlterarPublicacao(publicacaoID uint64, publicacao modelos.Publicacao) error {
	ps, erro := repositorio.db.Prepare("update publicacoes set titulo = ?, conteudo = ? where id = ?")
	if erro != nil {
		return erro
	}
	defer ps.Close()

	if _, erro := ps.Exec(publicacao.Titulo, publicacao.Conteudo, publicacaoID); erro != nil {
		return erro 
	}

	return nil
}