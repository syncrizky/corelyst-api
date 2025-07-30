package middleware

import (
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

var jwtKey = []byte("corelyst") // ganti sesuai secret production

// GenerateJWT membuat token dengan expiration 1 jam
func GenerateJWT(username string) (string, error) {
	claims := &jwt.RegisteredClaims{
		Subject:   username,
		ExpiresAt: jwt.NewNumericDate(time.Now().Add(1 * time.Hour)),
		IssuedAt:  jwt.NewNumericDate(time.Now()),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(jwtKey)
}

// JWTMiddleware memvalidasi token dari header Authorization
func JWTMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		authHeader := r.Header.Get("Authorization")
		if authHeader == "" {
			fmt.Println("🚫 Token kosong di Authorization header")
			http.Error(w, "Unauthorized: token tidak ada", http.StatusUnauthorized)
			return
		}

		if !strings.HasPrefix(authHeader, "Bearer ") {
			fmt.Println("🚫 Format header salah:", authHeader)
			http.Error(w, "Unauthorized: format token salah", http.StatusUnauthorized)
			return
		}

		tokenStr := strings.TrimPrefix(authHeader, "Bearer ")

		token, err := jwt.Parse(tokenStr, func(t *jwt.Token) (interface{}, error) {
			// Pastikan algoritma cocok
			if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
				fmt.Println("🚫 Metode tanda tangan tidak sesuai")
				return nil, fmt.Errorf("unexpected signing method: %v", t.Header["alg"])
			}
			return jwtKey, nil
		})

		if err != nil {
			fmt.Println("🚫 Gagal parse token:", err)
			http.Error(w, "Unauthorized: token invalid", http.StatusUnauthorized)
			return
		}

		if !token.Valid {
			fmt.Println("🚫 Token tidak valid")
			http.Error(w, "Unauthorized: token tidak valid", http.StatusUnauthorized)
			return
		}

		fmt.Println("✅ Token valid, akses diperbolehkan")
		next.ServeHTTP(w, r)
	})
}
