package sessao

import (
	"encoding/json"
	"net/http"
)

func Auth(response http.ResponseWriter, request *http.Request) {
	usuario, err := NewFromJson(request.Body)

	if err != nil {
		response.WriteHeader(http.StatusInternalServerError)
		response.Write([]byte(err.Error()))
		return
	}

	payload, _ := json.Marshal(usuario)

	response.Write(payload)
}
