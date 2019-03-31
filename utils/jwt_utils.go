package utils

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"time"
)

func GenToken(id int) (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)
	claims := make(jwt.MapClaims)
	claims["exp"] = time.Now().Add(time.Hour * time.Duration(10)).Unix()
	claims["iat"] = time.Now().Unix()
	claims["jti"] = id
	claims["iss"] = "www.allocmem.com"
	token.Claims = claims
	tokenString, err := token.SignedString([]byte("anteoy@gmail.com-secret"))
	if err != nil {
		fmt.Errorf("gen token 失败, err = %+v\n", err)
		return "", err
	}
	return tokenString, nil
}

func ValidateToken(tokenString string) bool {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// Don't forget to validate the alg is what you expect:
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok || jwt.SigningMethodHS256.Name != token.Header["alg"] {
			return nil, fmt.Errorf("Unexpected signing method: %v\n", token.Header["alg"])
		}
		return []byte("anteoy@gmail.com-secret"), nil
	})

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		fmt.Println(claims["jti"])
		return true
	} else {
		fmt.Println(err)
		return false
	}
}
