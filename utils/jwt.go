package utils

import (
    "fmt"
    "errors"
    "cloudSharePlatform/db"
    jwt "github.com/dgrijalva/jwt-go"
)

// 加密字符串
var JwtSecureSecret string = "secureSecretText"

// 结构
type JwtCustomClaims struct {
    LoginName string `json: "loginname"`
    jwt.StandardClaims
}

// 生成
func JwtCreate (loginName string) (string, error) {
    customClaims := JwtCustomClaims {
        LoginName: loginName,
        StandardClaims: jwt.StandardClaims {
            ExpiresAt: 15000,
            Issuer:    "nameOfWebsiteHere",
        },
    }
    
    token := jwt.NewWithClaims(jwt.SigningMethodHS256, customClaims)
    
    signedToken, err := token.SignedString([]byte(JwtSecureSecret))
    if err != nil {
        return "", err
    }
    return signedToken, nil
}

// 验证 -> 返回(loginname, error)
func JwtValidate (tokenStr string) (string, error) {
    customClaims := JwtCustomClaims{}
    token, err := jwt.ParseWithClaims(
    	tokenStr,
    	&customClaims,
    	func(token *jwt.Token) (interface{}, error) {
    		return []byte(JwtSecureSecret), nil 
    	},
    )
    if err != nil {
        return "", err
    }
    
    claims, ok := token.Claims.(*JwtCustomClaims)
    if !ok {
    	return "", errors.New("trans type nok")
    }
    
    // 数据库中查询用户登录账号
    loginname := claims.LoginName
    fmt.Println("loginName ============>", loginname)
    
    // 用获取用户信息
    dbResult, err := db.MongoFindOne("cloudshareplatform", "user", db.BsonM{ "loginname": loginname })
    if err != nil {
        return "", err
    }
    // 获取数据库查询结果
    var dbResultData db.MongoUser
    err = dbResult.Decode(&dbResultData)
    if err != nil {
        return "", err
    }
    
    return loginname, nil
}