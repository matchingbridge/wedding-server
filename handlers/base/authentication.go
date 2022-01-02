package base

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/dgrijalva/jwt-go"
	"github.com/dgrijalva/jwt-go/request"
)

func getReqIdFromToken(r *http.Request) (id string, err error) {
	token, claims, err := parseTokenClaims(r)
	if err != nil {
		return
	}
	if token.Valid {
		id = claims["sub"].(string)
		return
	} else {
		err = fmt.Errorf("forbidden no auth")
		return
	}
}

func parseTokenClaims(r *http.Request) (token *jwt.Token, claims jwt.MapClaims, err error) {
	tokenString, err := parseBearerToken(r)
	if err != nil {
		return
	}

	token, err = jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte("wedding"), nil
	})
	if err != nil {
		return
	}

	var ok bool
	if claims, ok = token.Claims.(jwt.MapClaims); ok {
		return
	} else {
		err = fmt.Errorf("jwt token forbidden")
		return
	}
}

func parseBearerToken(r *http.Request) (token string, err error) {
	const BearerPrefix = "Bearer "
	token, err = request.HeaderExtractor([]string{"Authorization"}).ExtractToken(r)
	if err != nil {
		return
	}
	return strings.TrimPrefix(token, BearerPrefix), nil
}
