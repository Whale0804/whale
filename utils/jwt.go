package utils

import (
	"errors"
	"fmt"
	"github.com/astaxie/beego"
	"github.com/dgrijalva/jwt-go"
	"time"
)

type WhaleClaims struct {
	Uid   int    `json: "uid"`
	Uname string `json: "uname"`
	jwt.StandardClaims
}

func ParaseToken(authorization string) (*WhaleClaims, bool, error) {
	token, _ := jwt.ParseWithClaims(authorization, &WhaleClaims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, nil
		}
		return []byte(beego.AppConfig.String("authKey")), nil
	})
	if claims, ok := token.Claims.(*WhaleClaims); ok && token.Valid {
		fmt.Println("claims:", claims)
		return claims, true, nil
	}
	return nil, false, errors.New("token invalid")
}

func GenToken(uid int, uname string) (string, error) {
	expireToken := time.Now().Add(time.Hour * 24).Unix()
	claims := WhaleClaims{
		uid,
		uname,
		jwt.StandardClaims{
			ExpiresAt: expireToken,
			Issuer:    "www.whale4cloud.com",
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(beego.AppConfig.String("authKey")))
}
