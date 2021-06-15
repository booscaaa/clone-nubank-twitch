package sessao

import "net/http"

func Refresh(response http.ResponseWriter, request *http.Request) {
	response.Write([]byte("Refresh"))
}
