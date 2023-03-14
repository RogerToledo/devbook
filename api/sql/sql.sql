CREATE DATABASE IF NOT EXISTS DEVBOOK;
USE DEVBOOK;

DROP TABLE usuarios;

CREATE TABLE usuarios(
    id int auto_increment primary key,
    nome varchar(50) not null,
    nick varchar(50) not null unique,
    email varchar(50) not null unique,
    senha varchar(60) not null unique,
    criadoEm timestamp default current_timestamp()
) ENGINE=INNODB;

CREATE TABLE seguidores (
  seguidor_id int NOT NULL,
  usuario_id int NOT NULL,
  PRIMARY KEY (seguidor_id,usuario_id),
  KEY seguidor_id (seguidor_id),
  KEY seguidores_ibfk_1 (usuario_id),
  CONSTRAINT seguidores_ibfk_1 FOREIGN KEY (usuario_id) REFERENCES usuarios (id) ON DELETE CASCADE,
  CONSTRAINT seguidores_ibfk_2 FOREIGN KEY (seguidor_id) REFERENCES usuarios (id) ON DELETE CASCADE
) ENGINE=InnoDB;