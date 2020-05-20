package common

import (
    "demo_items/gin_project/gin_vue_v2/model"
    "github.com/dgrijalva/jwt-go"
    "time"
)

var jwtKey = []byte("a_jwtKey")
type Claims struct {
    UserId uint
    jwt.StandardClaims
}

// 发放token函数
func ReleaseToken(user model.User) (string, error){
    // 过期时间
    expirationTime := time.Now().Add(7 * 24 * time.Hour)

    claims := &Claims{
        UserId: user.ID,
        StandardClaims: jwt.StandardClaims{
            ExpiresAt: expirationTime.Unix(),
            IssuedAt:  time.Now().Unix(),
            Issuer:    "xiaomaike",
            Subject:   "userToken",
        },
    }

    // 生成Token
    token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
    tokenString, err := token.SignedString(jwtKey)
        if err != nil {
        return "", err
    }
    return tokenString, nil
}

// 解析token函数
func ParseToken(tokenString string) (*jwt.Token, *Claims, error){
    claims := &Claims{}
    token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (i interface{}, err error){
        return jwtKey, err
    })
    return token, claims, err
}

