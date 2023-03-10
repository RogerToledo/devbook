CREATE DATABASE IF NOT EXISTS DEVBOOK;
USE DEVBOOK;

DROP TABLE seguidores;
DROP TABLE usuarios;

CREATE TABLE usuarios(
    id int auto_increment primary key,
    nome varchar(50) not null,
    nick varchar(50) not null unique,
    email varchar(50) not null unique,
    senha varchar(20) not null unique,
    criadoEm timestamp default current_timestamp()
) ENGINE=INNODB;

CREATE TABLE seguidores(
	usuario_id int not null,
	FOREIGN KEY (usuario_id)
	REFERENCES usuarios(id)
	ON DELETE CASCADE,	
	
	seguidor_id int not null,
	FOREIGN KEY (seguidor_id)
	REFERENCES usuarios(id)
	ON DELETE CASCADE,
	
	PRIMARY KEY(usuario_id, seguidor_id)
) ENGINE=INNODB;