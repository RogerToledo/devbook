package repositorios

import (
	"api/src/modelos"
	"database/sql"
	"errors"
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

func (repositorio *publicacoes) BuscarPorUsuario(ID uint64) ([]modelos.Publicacao, error) {
	linhas, erro := repositorio.db.Query(
		`select p.id, p.titulo, p.conteudo, p.curtidas, p.criadaEm, p.autor_id, u.nick 
		from publicacoes p
		inner join usuarios u on p.autor_id = u.id 
		where p.autor_id = ?
		order by p.criadaEm desc`, ID,
	)
	if erro != nil {
		return []modelos.Publicacao{}, erro
	}

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
		); erro != nil {
			return []modelos.Publicacao{}, erro
		} 

		publicacoes = append(publicacoes, publicacao)
	}

	return publicacoes, nil
}

func (repositorio *publicacoes) Alterar(publicacaoID uint64, publicacao modelos.Publicacao) error {
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

func (repositorio *publicacoes) Deletar(publicacaoID uint64) error {
	if _, erro := repositorio.db.Query("delete from publicacoes where id = ?", publicacaoID); erro != nil {
		return erro
	}

	return nil
}

func (repositorio *publicacoes) Curtir(publicacaoID, usuarioID uint64) error {
	existe, erro := repositorio.existeCurtida(publicacaoID, usuarioID)
	if erro != nil {
		return erro
	}

	if existe {
		repositorio.DeixarCurtir(publicacaoID, usuarioID)
		return nil
	}
	
	ps, erro := repositorio.db.Prepare("insert into curtidas (publicacao_id, usuario_id) values (?, ?)")
	if erro != nil {
		return erro
	}
	defer ps.Close()

	if _, erro := ps.Exec(publicacaoID, usuarioID); erro != nil {
		return erro
	}

	if erro := repositorio.atualizarCurtidasPublicacao(publicacaoID); erro != nil {
		return erro
	}

	return nil
}

func (repositorio *publicacoes) DeixarCurtir(publicaoID, usuarioID uint64) error {
	existe, erro := repositorio.existeCurtida(publicaoID, usuarioID)
	if erro != nil {
		return erro
	}

	if !existe {
		return errors.New("voce ainda não curtiu essa publicacao")
	}

	if erro := repositorio.deletaCurtida(publicaoID, usuarioID); erro != nil {
		return erro
	}

	if erro := repositorio.atualizarCurtidasPublicacao(publicaoID); erro != nil {
		return erro
	}

	return nil
}

func (repositorio *publicacoes) contarCurtidas(publicacaoID uint64) (uint64, error) {
	linha, erro := repositorio.db.Query("select count(1) from curtidas c where c.publicacao_id = ?", publicacaoID)
	if erro != nil {
		return 0, erro
	}
	defer linha.Close()

	var modelo modelos.Publicacao
	if linha.Next() {
		if linha.Scan(
			&modelo.Curtidas,
		); erro != nil {
			return 0, erro
		}
	}

	curtidas := modelo.Curtidas

	return curtidas, nil 
}

func (repositorio *publicacoes) atualizarCurtidasPublicacao(publicacaoID uint64) error {
	curtidas, erro := repositorio.contarCurtidas(publicacaoID)
	if erro != nil {
		return erro
	}

	ps, erro := repositorio.db.Prepare("update publicacoes set curtidas = ? where id = ?")
	if erro != nil {
		return erro
	}
	defer ps.Close()

	if _, erro := ps.Exec(curtidas, publicacaoID); erro != nil {
		return erro
	}

	return nil
}

func (repositorio *publicacoes) existeCurtida(publicacaoID, usuarioID uint64) (bool, error) {
	linha, erro := repositorio.db.Query(
		`select publicacao_id 
		from curtidas c 
		where c.publicacao_id = ? and c.usuario_id = ?`, publicacaoID, usuarioID,
	)
	if erro != nil {
		return false, erro
	}

	var curtida modelos.Curtida
	if linha.Next() {
		if erro := linha.Scan(
			&curtida.PublicacaoID,
		); erro != nil {
			return false, erro
		}
	}

	if curtida.PublicacaoID == 0 {
		return false, nil
	}

	return true, nil
}

func (repositorio *publicacoes) deletaCurtida(publicacaoID, usuarioID uint64) error {
	if _, erro := repositorio.db.Query(
		"delete from curtidas where publicacao_id = ? and usuario_id = ?", publicacaoID, usuarioID,
	); erro != nil {
		return erro
	}

	return nil
}
