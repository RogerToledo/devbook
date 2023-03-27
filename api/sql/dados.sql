INSERT INTO usuarios (nome, nick, email, senha) VALUES -- senha 123
("Usuario1", "user1", "usuario1@mail.com", "$2a$10$NilCI7IgxlA17BTcB.XC8ePuwVeINybhvoAwoTdoU7.eDLRzbjfHK"),
("Usuario2", "user2", "usuario2@mail.com", "$2a$10$NilCI7IgxlA17BTcB.XC8ePuwVeINybhvoAwoTdoU7.eDLRzbjfHK"),
("Usuario3", "user3", "usuario3@mail.com", "$2a$10$NilCI7IgxlA17BTcB.XC8ePuwVeINybhvoAwoTdoU7.eDLRzbjfHK"),
("Usuario4", "user4", "usuario4@mail.com", "$2a$10$NilCI7IgxlA17BTcB.XC8ePuwVeINybhvoAwoTdoU7.eDLRzbjfHK"),
("Usuario5", "user5", "usuario5@mail.com", "$2a$10$NilCI7IgxlA17BTcB.XC8ePuwVeINybhvoAwoTdoU7.eDLRzbjfHK");

INSERT INTO seguidores (seguidor_id, usuario_id) VALUES
(1, 2), (1, 3), (2, 3), (2, 4), (2, 5), (3, 5), (3, 1),
(3, 2), (4, 1), (4, 3), (5, 3),

INSERT INTO publicacoes
(titulo, conteudo, autor_id, autor_nick, curtidas, criadaEm) VALUES
('1_Publicacao-1', '1_Publicacao-1 ...', 1, 'user1' , 0, now()),
('2_Publicacao-1', '2_Publicacao-1 ...', 2, 'user2' , 0, now()),
('3_Publicacao-1', '3_Publicacao-1 ...', 3, 'user3' , 0, now()),
('4_Publicacao-1', '4_Publicacao-1 ...', 4, 'user4' , 0, now()),
('5_Publicacao-1', '5_Publicacao-1 ...', 5, 'user5' , 0, now()),




