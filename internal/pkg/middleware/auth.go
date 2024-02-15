package middleware

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/golang-jwt/jwt"
)

var roleMap = map[string]int{
	"Admin":    1,
	"Employee": 0,
}

func verifyToken(tokenString string) (jwt.Claims, error) {
	signingKey := []byte("keymaker")
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return signingKey, nil
	})
	if err != nil {
		return nil, err
	}
	return token.Claims, err
}

func AuthMiddleware(handler http.Handler, role string) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		fmt.Println("Role from tokne : ")
		fmt.Println("Role  : ")

		tokenString := r.Header.Get("Authorization")
		if len(tokenString) == 0 {
			w.WriteHeader(http.StatusUnauthorized)
			w.Write([]byte("Missing Authorization Header"))
			return
		}
		tokenString = strings.Replace(tokenString, "Bearer ", "", 1)
		claims, err := verifyToken(tokenString)
		if err != nil {
			w.WriteHeader(http.StatusUnauthorized)
			w.Write([]byte("Error verifying JWT token: " + err.Error()))
			return
		}
		roleFromToken := claims.(jwt.MapClaims)["role_id"].(float64)

		// r.Header.Set("role", role)
		roleVal := roleMap[role]

		fmt.Println("Role from tokne : ", roleFromToken)
		fmt.Println("Role  : ", roleVal)
		if roleVal != int(roleFromToken) {
			http.Error(w, "Unauthorized : ", http.StatusUnauthorized)
			return
		}

		handler.ServeHTTP(w, r)
	})
}
