package middleware

import "net/http"

func VerifyAuth(next http.Handler) http.Handler {
	return http.HandlerFunc(func(response http.ResponseWriter, request *http.Request) {
		if request.Method != "OPTIONS" {
			bearerToken := request.Header.Get("Authorization")

			if isAuth, usuario := VerifyToken(bearerToken); isAuth {
				request = SetContextData(request, usuario)
				next.ServeHTTP(response, request)
				return
			} else {
				response.WriteHeader(http.StatusUnauthorized)
			}
		}

		next.ServeHTTP(response, request)
	})
}
