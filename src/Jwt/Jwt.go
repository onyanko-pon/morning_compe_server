package Jwt

import (
	jwt "github.com/dgrijalva/jwt-go"
	"os"
)

type Claims struct {
    Sub string
    Exp int
}

func GenerateJwtToken(c *Claims) string {
	tokenWithHeader := jwt.New(jwt.SigningMethodHS256)
	claims := tokenWithHeader.Claims.(jwt.MapClaims)
	claims["sub"] = c.Sub
	claims["exp"] = c.Exp
	// 電子署名
	tokenString, _ := tokenWithHeader.SignedString([]byte(os.Getenv("JWT_SECRET_KEY")))
	return tokenString
}

func TokenToClaims(tokenstring string) *Claims {
	token, _ := jwt.Parse(tokenstring, func(token *jwt.Token) (interface{}, error) {
		return []byte("foobar"), nil
	})
	claims_map := token.Claims.(jwt.MapClaims)
	return &Claims{
		Sub: claims_map["sub"].(string),
		Exp: int(claims_map["exp"].(float64)),
	}
}

func Valid(tokenstring string) bool {
	token, _ := jwt.Parse(tokenstring, func(token *jwt.Token) (interface{}, error) {
		return []byte(os.Getenv("JWT_SECRET_KEY")), nil
	})
	return token.Valid
}