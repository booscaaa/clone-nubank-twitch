package middleware

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/dgrijalva/jwt-go"
	"github.com/mitchellh/mapstructure"
)

type Auth struct {
	Token   string `json:"token"`
	Refresh string `json:"refresh"`
	Type    string `json:"type"`
}

type TokenAuthJwt struct {
	Usuario Usuario `json:"usuario"`
	Exp     int64   `json:"exp"`
	jwt.StandardClaims
}

func ExtractToken(bearerToken string) string {
	strArr := strings.Split(bearerToken, " ")

	if len(strArr) == 2 {
		return strArr[1]
	}

	return ""
}

func CreateToken(tokenAuthJwt TokenAuthJwt, hash string) Auth {
	token := jwt.NewWithClaims(jwt.GetSigningMethod("HS256"), &tokenAuthJwt)

	hashEnv := os.Getenv("BCRYPT_HASH")
	tokenString, err := token.SignedString([]byte(hashEnv))
	if err != nil {
		log.Fatalln(err)
	}

	return Auth{
		Token:   tokenString,
		Refresh: hash,
		Type:    "refreshToken",
	}
}

func VerifyToken(bearerToken string) (bool, Usuario) {
	usuario := Usuario{}

	token, err := jwt.Parse(ExtractToken(bearerToken), func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Falha na validação do token de acesso!")
		}

		hashEnv := os.Getenv("BCRYPT_HASH")
		return []byte(hashEnv), nil
	})

	if err != nil {
		return false, usuario
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		mapstructure.Decode(claims["usuario"], &usuario)
	} else {
		return false, usuario
	}

	return true, usuario
}

func SetContextData(r *http.Request, u *Usuario) (ro *http.Request) {
	ctx := r.Context()
	ctx = context.WithValue(ctx, 1, u)
	ro = r.WithContext(ctx)
	return
}

func GetContextData(r *http.Request) (u Usuario) {
	d := *r.Context().Value(1).(*Usuario)
	return d
}
