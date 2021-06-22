package sessao

import (
	"api/factory"
	"api/middleware"
	"encoding/json"
	"net/http"
	"strings"
	"time"
)

func Refresh(response http.ResponseWriter, request *http.Request) {
	bearerToken := request.Header.Get("Authorization")
	bearTokenSlice := strings.Split(bearerToken, " ")

	if len(bearTokenSlice) < 4 {
		response.WriteHeader(http.StatusUnauthorized)
		response.Write([]byte("Token inválido ou não informado"))
		return
	}

	refreshToken := bearTokenSlice[2]
	typeToken := bearTokenSlice[3]
	var revogado bool
	var idUsuario int
	var usuario middleware.Usuario

	db := factory.GetConnection()
	defer db.Close()

	{
		rows, err := db.Query(`
			SELECT revogado, id_usuario FROM auth where 
			refresh_token = $1 and
			tipo_token = $2;
		`, refreshToken, typeToken)

		if err != nil {
			response.WriteHeader(http.StatusInternalServerError)
			response.Write([]byte(err.Error()))
			return
		}

		for rows.Next() {
			rows.Scan(&revogado, &idUsuario)
		}

		if revogado {
			response.WriteHeader(http.StatusUnauthorized)
			response.Write([]byte("Token revogado"))
			return
		}
	}

	{
		rows, err := db.Query(`
			SELECT * FROM usuario where id = $1;
		`,
			idUsuario,
		)

		if err != nil {
			response.WriteHeader(http.StatusInternalServerError)
			response.Write([]byte("Algum problema ocorreu durante a validação do usuário"))
			return
		}

		for rows.Next() {
			var id int
			var login string
			var senha string
			var email string
			var dataCriacao string
			var idGrupoUsuario int

			rows.Scan(&id, &login, &senha, &email, &dataCriacao, &idGrupoUsuario)

			usuarioMiddleware, err := middleware.NewUsuario(
				id,
				login,
				senha,
				email,
				dataCriacao,
				idGrupoUsuario,
			)

			if err != nil {
				response.WriteHeader(http.StatusInternalServerError)
				response.Write([]byte(err.Error()))
				return
			}

			usuario = *usuarioMiddleware
		}
	}

	if usuario.ID == 0 {
		response.WriteHeader(http.StatusUnauthorized)
		response.Write([]byte("Usuário não existe na base de dados"))
		return
	}

	tokenAuthJwt := middleware.TokenAuthJwt{
		Usuario: usuario,
		Exp:     time.Now().Add(time.Minute * 5).Unix(),
	}

	token := middleware.CreateToken(tokenAuthJwt, refreshToken)
	payload, _ := json.Marshal(token)

	response.Write([]byte(payload))
}
