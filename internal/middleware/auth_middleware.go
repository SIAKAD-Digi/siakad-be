package middleware

import (
	"fmt"
	"os"
	"siakad-digi/internal/exception"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

func JWTMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		authHeader := ctx.GetHeader("Authorization")
		if authHeader == "" || !strings.HasPrefix(authHeader, "Bearer ") {
			panic(exception.NewUnauthorizeError("token tidak ditemukan"))
		}

		tokenStr := strings.TrimPrefix(authHeader, "Bearer ")

		secretKey := os.Getenv("JWT_SECRET_KEY")

		token, err := jwt.Parse(tokenStr, func(token *jwt.Token) (any, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, jwt.ErrTokenSignatureInvalid
			}
			return []byte(secretKey), nil
		})

		if err != nil {
			fmt.Println(err, "jwt errooo")
			panic(exception.NewUnauthorizeError("token expired"))
		}

		if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
			ctx.Set("userClaims", claims)
			ctx.Next()
		} else {
			panic(exception.NewUnauthorizeError("token expired"))
		}
	}
}
