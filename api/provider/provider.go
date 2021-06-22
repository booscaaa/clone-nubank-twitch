package provider

import (
	"api/controller/cartao_credito"
	"api/controller/sessao"
	"api/middleware"
	"net/http"

	"github.com/gorilla/mux"
)

func GetProvider() *mux.Router {
	r := mux.NewRouter()

	r.HandleFunc("/sessao", sessao.Auth).Methods("POST", "OPTIONS")
	r.HandleFunc("/sessao", sessao.Refresh).Methods("GET", "OPTIONS")

	r.Handle("/compras/credito", middleware.VerifyAuth(http.HandlerFunc(cartao_credito.Get))).Methods("GET", "OPTIONS")

	r.HandleFunc("/", func(response http.ResponseWriter, request *http.Request) {
		response.Write([]byte("funcionando"))
	}).Methods("GET", "OPTIONS")

	return r
}
