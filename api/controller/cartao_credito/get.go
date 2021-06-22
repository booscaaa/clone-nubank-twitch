package cartao_credito

import (
	"api/factory"
	"api/middleware"
	"encoding/json"
	"net/http"
)

func Get(response http.ResponseWriter, request *http.Request) {
	usuario := middleware.GetContextData(request)
	var cartoes []Cartao

	db := factory.GetConnection()

	defer db.Close()

	{
		rows, err := db.Query(`
			SELECT cc.*, u.login, u.email FROM cartao_credito cc
			INNER JOIN usuario u ON u.id = cc.id_usuario 
			WHERE id_usuario = $1;
		`, usuario.ID)

		if err != nil {
			response.WriteHeader(http.StatusInternalServerError)
			response.Write([]byte(err.Error()))
			return
		}

		for rows.Next() {
			var id int
			var valor float64
			var descricao string
			var local string
			var idUsuario int
			var login string
			var email string

			rows.Scan(&id, &valor, &descricao, &local, &idUsuario, &login, &email)

			cartao, err := New(id, valor, descricao, local, idUsuario, login, email)

			if err != nil {
				response.WriteHeader(http.StatusInternalServerError)
				response.Write([]byte(err.Error()))
			}

			cartoes = append(cartoes, *cartao)
		}
	}

	payload, _ := json.Marshal(cartoes)
	response.Write(payload)
}
