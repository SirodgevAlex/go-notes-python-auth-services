package auth

import (
	"errors"
	"fmt"
	"net/http"
	"os"
    "strconv"

	"github.com/dgrijalva/jwt-go"
)

var jwtKey = []byte(os.Getenv("JWT_SECRET_KEY"))

type Claims struct {
    UserID int `json:"sub"`
    jwt.StandardClaims
}

func ValidateToken(next http.HandlerFunc) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		tokenString := r.Header.Get("Authorization")
		if tokenString == "" {
			http.Error(w, "Unauthorized: missing token", http.StatusUnauthorized)
			return
		}

		token, err := jwt.ParseWithClaims(tokenString, &Claims{}, func(token *jwt.Token) (interface{}, error) {
			return jwtKey, nil
		})
		if err != nil || !token.Valid {
			http.Error(w, "Unauthorized: invalid token", http.StatusUnauthorized)
			return
		}

		next(w, r)
	})
}

func GetUserIdFromToken(tokenString string) (int, error) {
    token, _, err := new(jwt.Parser).ParseUnverified(tokenString, &Claims{})
    if err != nil {
        fmt.Println("Error during ParseUnverified:", err)
        return 0, fmt.Errorf("failed to parse token: %s", err)
    }

    claims, ok := token.Claims.(*Claims)
    if !ok {
        return 0, errors.New("invalid claims")
    }

    userID, err := strconv.Atoi(fmt.Sprintf("%v", claims.UserID))
    if err != nil {
        return 0, fmt.Errorf("failed to parse user ID: %s", err)
    }

    return userID, nil
}