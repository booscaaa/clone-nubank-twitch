package middleware

import "net/http"

func Cors(next http.Handler) http.Handler {
	return http.HandlerFunc(func(response http.ResponseWriter, request *http.Request) {
		response.Header().Set("Access-Control-Allow-Methods", "POST, GET, PUT, DELETE, OPTIONS")
		response.Header().Set("Content-Type", "application/json")
		response.Header().Set("Access-Control-Allow-Origin", "*")
		response.Header().Set("Acces-Control-Allow-Headers", "Content-Type, Authorization")
		if request.Method == "OPTIONS" {
			response.WriteHeader(http.StatusOK)
		}

		next.ServeHTTP(response, request)
	})
}
