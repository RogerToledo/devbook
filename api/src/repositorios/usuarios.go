package repositorios

import (
	"api/src/modelos"
	"database/sql"
)

type usuarios struct {
	db *sql.DB
}

func NovoRepositorioDeUsuarios(db *sql.DB) *usuarios {
	return &usuarios{db}
}

func (repositorio usuarios) Criar(usuario modelos.Usuario) (uint64, error) {
	ps, erro := repositorio.db.Prepare(
		"insert into usuarios (nome, nick, email, senha) values (?, ?, ?, ?)",
	)
	if erro != nil {
		return 0, erro
	}
	defer ps.Close()

	resultado, erro := ps.Exec(usuario.Nome, usuario.Nick, usuario.Email, usuario.Senha)
	if erro != nil {
		return 0, erro
	}

	ultimoId, erro := resultado.LastInsertId()
	if erro != nil {
		return 0, erro
	}

	return uint64(ultimoId), nil
}

func (repositorio usuarios) Buscar() ([]modelos.Usuario, error) {
	linhas, erro := repositorio.db.Query(
		"select id, nome, nick, email, criadoEm from usuarios",
	)
	if erro != nil {
		return nil, erro
	}
	defer linhas.Close()

	var usuarios []modelos.Usuario

	for linhas.Next() {
		var usuario modelos.Usuario

		if erro = linhas.Scan(
			&usuario.ID,
			&usuario.Nome,
			&usuario.Nick,
			&usuario.Email,
			&usuario.CriadoEm,
		); erro != nil {
			return nil, erro
		}

		usuarios = append(usuarios, usuario)
	}

	return usuarios, nil
}

func (repositorio usuarios) BuscarPorId(id uint64) (modelos.Usuario, error) {
	linhas, erro := repositorio.db.Query("select id, nome, nick, email, criadoEm from usuarios where id = ?", id)
	if erro != nil {
		return modelos.Usuario{}, erro
	}
	defer linhas.Close()

	var usuario modelos.Usuario

	if linhas.Next() {
		if erro = linhas.Scan(
			&usuario.ID,
			&usuario.Nome,
			&usuario.Nick,
			&usuario.Email,
			&usuario.CriadoEm,
		); erro != nil {
			return modelos.Usuario{}, erro
		}
	}

	return usuario, nil
}

func (repositorio usuarios) BuscarPorEmail(email string) (modelos.Usuario, error) {
	linha, erro := repositorio.db.Query("select id, senha from usuarios where email = ?", email)
	if erro != nil {
		return modelos.Usuario{}, erro
	}
	defer linha.Close()

	var usuario modelos.Usuario

	if linha.Next() {
		if erro := linha.Scan(
			&usuario.ID,
			&usuario.Senha,
		); erro != nil {
			return modelos.Usuario{}, erro
		}
	}
	return usuario, nil
}

func (repositorio usuarios) Atualizar(ID uint64, usuario modelos.Usuario) error {
	ps, erro := repositorio.db.Prepare(
		"update usuarios set nome = ?, nick = ?, email = ? where id = ?",
	)
	if erro != nil {
		return erro
	}
	defer ps.Close()

	if _, erro := ps.Exec(usuario.Nome, usuario.Nick, usuario.Email, ID); erro != nil {
		return erro
	}

	return nil
}

func (repositorio usuarios) Deletar(ID uint64) error {
	ps, erro := repositorio.db.Prepare("delete from usuarios where id = ?")
	if erro != nil {
		return erro
	}
	defer ps.Close()

	if _, erro := ps.Exec(ID); erro != nil {
		return erro
	}

	return nil
}

func (repositorio usuarios) Seguir(seguidorID, ID uint64) error {
	ps, erro := repositorio.db.Prepare("insert ignore into seguidores (seguidor_id, usuario_id) values (?, ?)")
	if erro != nil {
		return erro
	}
	defer ps.Close()

	if _, erro := ps.Exec(seguidorID, ID); erro != nil {
		return erro
	}

	return nil
}

func (repositorio usuarios) PararSeguirUsuario(seguidorID, ID uint64) error {
	ps, erro := repositorio.db.Prepare("delete ignore from seguidores where usuario_id = ? and seguidor_id = ?")
	if erro != nil {
		return erro
	}
	defer ps.Close()

	if _, erro := ps.Exec(seguidorID, ID); erro != nil {
		return erro
	}

	return nil
}

func (repositorio usuarios) ListarSeguindoSeguidores(ID uint64, tipo string) ([]modelos.Usuario, error) {
	var query string
	seguidor := "select u.nome, u.nick from seguidores s inner join usuarios u on s.usuario_id = u.id where s.seguidor_id = ?"
	seguindo := "select u.nome, u.nick from seguidores s inner join usuarios u on s.seguidor_id = u.id where s.usuario_id = ?"

	if tipo == "listar-seguidores" {
		query = seguidor
	} else if tipo == "listar-seguindo" {
		query = seguindo
	}

	linhas, erro := repositorio.db.Query(query, ID)
	if erro != nil {
		return nil, erro
	}
	defer linhas.Close()

	var usuarios []modelos.Usuario

	for linhas.Next() {
		var usuario modelos.Usuario

		if erro := linhas.Scan(
			&usuario.Nome,
			&usuario.Nick,
		); erro != nil {
			return nil, erro
		}

		usuarios = append(usuarios, usuario)
	}

	return usuarios, nil
}

func (repositorio usuarios) BuscarSenha(ID uint64) (string, error) {
	linha, erro := repositorio.db.Query("select senha from usuarios where id = ?", ID)
	if erro != nil {
		return "", erro
	}
	defer linha.Close()

	var usuario modelos.Usuario

	if linha.Next() {
		if erro := linha.Scan(&usuario.Senha); erro != nil {
			return "", erro
		}
	}

	return usuario.Senha, nil
}

func (repositorio usuarios) AlterarSenha(ID uint64, senha string) error {
	ps, erro := repositorio.db.Prepare("update usuarios set senha = ? where id = ?")
	if erro != nil {
		return erro
	}
	defer ps.Close()

	if _, erro := ps.Exec(senha, ID); erro != nil {
		return erro
	}

	return nil
}
