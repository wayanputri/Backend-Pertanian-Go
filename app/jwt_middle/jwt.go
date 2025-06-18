package jwt_middle

import (
	"context"
	"fmt"
	"strings"
	"time"

	"github.com/dgrijalva/jwt-go"
	"google.golang.org/grpc/metadata"
)

// CreateToken creates a JWT token based on the userId
func CreateToken(userId int, role string) (string, error) {
	claims := jwt.MapClaims{}
	claims["authorized"] = true
	claims["userId"] = userId
	claims["role"] = role
	claims["exp"] = time.Now().Add(time.Hour * 48).Unix()

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte("secret"))
}

// ExtractTokenUserId akan mengekstrak userId dari context.Context (gRPC)
func ExtractTokenUserId(ctx context.Context) (int, string, error) {
	// Ambil metadata dari context
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return 0, "", fmt.Errorf("no metadata in context")
	}

	// Ambil token dari metadata (misalnya dari header Authorization)
	tokens := md["authorization"]
	if len(tokens) == 0 {
		return 0, "", fmt.Errorf("authorization token not found")
	}

	// Ambil token dari header (berformat Bearer <token>)
	tokenString := tokens[0]
	tokenString = strings.TrimPrefix(tokenString, "Bearer ")

	// Parsing token
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// Pastikan token menggunakan metode signing yang benar
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method")
		}
		// Kembalikan kunci signing
		return []byte("secret"), nil
	})
	if err != nil || !token.Valid {
		return 0, "", fmt.Errorf("invalid token")
	}

	// Ambil claims dari token
	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return 0, "", fmt.Errorf("invalid claims")
	}

	// Ambil userId dari claims
	userId, ok := claims["userId"].(float64)
	if !ok {
		return 0, "", fmt.Errorf("userId not found in token")
	}

	role, ok := claims["role"].(string)
	if !ok {
		return 0, "", fmt.Errorf("role not found in token")
	}

	return int(userId), role, nil
}
