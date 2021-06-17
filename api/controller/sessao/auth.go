package sessao

import (
	"api/factory"
	"api/middleware"
	"encoding/json"
	"net/http"
	"os"
	"time"

	_ "github.com/lib/pq"
	"golang.org/x/crypto/bcrypt"
)

func Auth(response http.ResponseWriter, request *http.Request) {
	usuarioRequest, err := middleware.NewFromJson(request.Body)
	var usuario middleware.Usuario

	if err != nil {
		response.WriteHeader(http.StatusInternalServerError)
		response.Write([]byte(err.Error()))
		return
	}

	db := factory.GetConnection()
	defer db.Close()

	{
		rows, err := db.Query(`
			select * from usuario 
			where login = $1 and 
			senha = crypt($2, senha);
		`,
			usuarioRequest.Login,
			usuarioRequest.Senha,
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
		response.Write([]byte("Usuário ou senha inválidos"))
		return
	}

	tokenAuthJwt := middleware.TokenAuthJwt{
		Usuario: usuario,
		Exp:     time.Now().Add(time.Minute * 5).Unix(),
	}

	hashBcrypt := os.Getenv("BCRYPT_HASH")

	hashB, err := bcrypt.GenerateFromPassword([]byte(hashBcrypt), bcrypt.DefaultCost)

	if err != nil {
		response.WriteHeader(http.StatusInternalServerError)
		response.Write([]byte("Problema ao gerar o hash de bcrypt"))
		return
	}

	token := middleware.CreateToken(tokenAuthJwt, string(hashB))

	payload, _ := json.Marshal(token)

	response.Write(payload)
}
