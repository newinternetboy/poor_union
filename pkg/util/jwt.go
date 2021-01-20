/*
 * @Description:
 * @Author: Caoxiang
 * @Date: 2021-01-19 16:12:50
 * @LastEditors: Caoxiang
 */
package util

import (
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/newinternetboy/poor_union/pkg/setting"
)

var jwtSecret = []byte(setting.JwtSecret)

type Claims struct {
	UID int `json:"u_id"`
	jwt.StandardClaims
}

//生成token
func GenerateToken(uid int) (string, error) {
	nowTime := time.Now()
	expireTime := nowTime.Add(3 * time.Hour)

	claims := Claims{
		uid,
		jwt.StandardClaims{
			ExpiresAt: expireTime.Unix(),
			Issuer:    "gin-blog",
		},
	}

	tokenClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	token, err := tokenClaims.SignedString(jwtSecret)
	return token, err
}

//解析token
func ParseToken(token string) (*Claims, error) {
	tokenClaims, err := jwt.ParseWithClaims(token, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return jwtSecret, nil
	})

	if tokenClaims != nil {
		if claims, ok := tokenClaims.Claims.(*Claims); ok && tokenClaims.Valid {
			return claims, nil
		}
	}
	return nil, err
}
