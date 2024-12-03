package main

import (
	"fmt"
	"time"

	"github.com/dgrijalva/jwt-go"
)

func main() {
	jwtKey := []byte("skdjfskdfjsdkjfsdk")

	newToken := jwt.New(jwt.SigningMethodHS256)
	claims := newToken.Claims.(jwt.MapClaims)
	claims["user"] = "Charles Usario"
	claims["exp"] = time.Now().Add(time.Hour * 24).Unix()

	token,err := newToken.SignedString(jwtKey)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(token)

}
