package middleware

import (
	"encoding/json"
	"fmt"
	"io"
)

type Usuario struct {
	ID             int    `json:"id"`
	Login          string `json:"login"`
	Senha          string `json:"senha"`
	Email          string `json:"email"`
	DataCriacao    string `json:"dataCriacao"`
	IDGrupoUsuario int    `json:"idGrupoUsuario"`
}

func (u Usuario) isValid() error {
	if u.ID == 0 {
		return fmt.Errorf("Preencha o id do usuário")
	}

	if u.Login == "" {
		return fmt.Errorf("Preencha o login do usuário")
	}

	if u.Senha == "" {
		return fmt.Errorf("Preencha o senha do usuário")
	}

	if u.Email == "" {
		return fmt.Errorf("Preencha o email do usuário")
	}

	if u.IDGrupoUsuario == 0 {
		return fmt.Errorf("Preencha o id do grupo do usuário")
	}

	return nil
}

func (u Usuario) isValidLogin() error {
	if u.Login == "" {
		return fmt.Errorf("Preencha o login do usuário")
	}

	if u.Senha == "" {
		return fmt.Errorf("Preencha o senha do usuário")
	}

	return nil
}

func NewUsuario(id int, login string, senha string, email string, dataCriacao string, idGrupoUsuario int) (*Usuario, error) {
	usuario := Usuario{
		ID:             id,
		Login:          login,
		Senha:          senha,
		Email:          email,
		DataCriacao:    dataCriacao,
		IDGrupoUsuario: idGrupoUsuario,
	}

	if err := usuario.isValid(); err != nil {
		return nil, err
	}

	return &usuario, nil
}

func NewFromJson(body io.ReadCloser) (*Usuario, error) {
	var usuario Usuario

	if err := json.NewDecoder(body).Decode(&usuario); err != nil {
		return nil, err
	}

	if err := usuario.isValidLogin(); err != nil {
		return nil, err
	}

	return &usuario, nil
}
