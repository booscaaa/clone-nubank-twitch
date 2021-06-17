CREATE EXTENSION pgcrypto; 

CREATE TABLE acao(
    id SERIAL PRIMARY KEY NOT NULL,
    nome VARCHAR(40) NOT NULL
);

CREATE TABLE permissao(
    id SERIAL PRIMARY KEY NOT NULL,
    nome VARCHAR(40) NOT NULL
);

CREATE TABLE grupo_usuario(
    id SERIAL PRIMARY KEY NOT NULL,
    nome VARCHAR(120) NOT NULL
);

CREATE TABLE grupo_usuario_permissao(
    id SERIAL PRIMARY KEY NOT NULL,
    id_grupo_usuario INTEGER NOT NULL,
    id_permissao INTEGER NOT NULL,
    id_acao INTEGER NOT NULL,

    CONSTRAINT id_grupo_usuario_fk FOREIGN KEY(id_grupo_usuario) REFERENCES grupo_usuario(id),
    CONSTRAINT id_permissao_fk FOREIGN KEY(id_permissao) REFERENCES permissao(id),
    CONSTRAINT id_acao_fk FOREIGN KEY(id_acao) REFERENCES acao(id)
);

CREATE TABLE usuario(
    id SERIAL PRIMARY KEY NOT NULL,
    login VARCHAR(60) NOT NULL,
    senha VARCHAR(200) NOT NULL,
    email VARCHAR(100) NOT NULL,
    data_criacao TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    id_grupo_usuario INTEGER NOT NULL,

    CONSTRAINT id_grupo_usuario_fk FOREIGN KEY(id_grupo_usuario) REFERENCES grupo_usuario(id)
);

CREATE TABLE auth(
    id SERIAL PRIMARY KEY NOT NULL,
    tipo_token VARCHAR(20) NOT NULL,
    revogado boolean NOT NULL DEFAULT false,
    refresh_token VARCHAR(200) NOT NULL
);

CREATE TABLE cartao_credito(
    id SERIAL PRIMARY KEY NOT NULL,
    valor DECIMAL NOT NULL,
    descricao VARCHAR(200),
    local VARCHAR(200) NOT NULL,
    id_usuario INTEGER NOT NULL,

    CONSTRAINT id_usuario_fk FOREIGN KEY(id_usuario) REFERENCES usuario(id)
);

CREATE TABLE cartao_debito(
    id SERIAL PRIMARY KEY NOT NULL,
    valor DECIMAL NOT NULL,
    descricao VARCHAR(200),
    id_usuario INTEGER NOT NULL,

    CONSTRAINT id_usuario_fk FOREIGN KEY(id_usuario) REFERENCES usuario(id)
);