package cartao_credito

import "fmt"

type Cartao struct {
	ID        int     `json:"id"`
	Valor     float64 `json:"valor"`
	Descricao string  `json:"descricao"`
	Local     string  `json:"local"`
	Usuario   Usuario `json:"usuario"`
}

type Usuario struct {
	ID    int    `json:"id"`
	Login string `json:"login"`
	Email string `json:"email"`
}

func New(id int, valor float64, descricao string, local string, idUsuario int, login string, email string) (*Cartao, error) {
	cartao := &Cartao{
		ID:        id,
		Valor:     valor,
		Descricao: descricao,
		Local:     local,
		Usuario: Usuario{
			ID:    id,
			Login: login,
			Email: email,
		},
	}

	if isValid, err := cartao.isValid(); !isValid {
		return nil, err
	}

	return cartao, nil
}

func (cartao Cartao) isValid() (bool, error) {
	if cartao.ID == 0 {
		return false, fmt.Errorf("Preencha o id do cartão")
	}

	if cartao.Valor == 0.0 {
		return false, fmt.Errorf("Preencha o valor do cartão")
	}

	if cartao.Valor < 0.0 {
		return false, fmt.Errorf("O valor da compra não pode ser menor que zero")
	}

	if cartao.Descricao == "" {
		return false, fmt.Errorf("Preencha a descrição")
	}

	if cartao.Local == "" {
		return false, fmt.Errorf("Preencha o local")
	}

	if cartao.Usuario.ID == 0 {
		return false, fmt.Errorf("Preencha o id do usuário")
	}

	return true, nil
}
