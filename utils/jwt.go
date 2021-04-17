package utils

import (
	"cloudSharePlatform/db"
	"errors"
	"fmt"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
	primitive "go.mongodb.org/mongo-driver/bson/primitive"
)

// 加密字符串
var JwtSecureSecret string = "secureSecretText"

// 过期时间
var JwtExpiresAt int64 = int64(time.Now().Unix() + (3600 * 24 * 365))

// 结构
type JwtCustomClaims struct {
	LoginName string `json: "loginname"`
	jwt.StandardClaims
}

// 生成
func JwtCreate(loginName string) (string, error) {
	customClaims := JwtCustomClaims{
		LoginName: loginName,
		StandardClaims: jwt.StandardClaims{
			NotBefore: int64(time.Now().Unix()), //签名生效时间
			ExpiresAt: JwtExpiresAt,             // 过期时间 (一天)
			Issuer:    "cloudSharePlatform",     //签名发行者
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, customClaims)

	signedToken, err := token.SignedString([]byte(JwtSecureSecret))
	if err != nil {
		return "", err
	}
	return signedToken, nil
}

// 验证 -> 返回(_id, error)
func JwtValidate(tokenStr string) (primitive.ObjectID, error) {
	customClaims := JwtCustomClaims{}
	token, err := jwt.ParseWithClaims(
		tokenStr,
		&customClaims,
		func(token *jwt.Token) (interface{}, error) {
			return []byte(JwtSecureSecret), nil
		},
	)
	if err != nil {
		return primitive.NilObjectID, err
	}

	claims, ok := token.Claims.(*JwtCustomClaims)
	if !ok {
		return primitive.NilObjectID, errors.New("trans type nok")
	}

	// 数据库中查询用户登录账号
	loginname := claims.LoginName

	// 用获取用户信息
	dbResult, err := db.MongoFindOne("cloudshareplatform", "user", db.BsonM{"loginname": loginname})
	if err != nil {
		return primitive.NilObjectID, err
	}
	// 获取数据库查询结果
	var mongoUser db.MongoUser
	err = dbResult.Decode(&mongoUser)
	if err != nil {
		return primitive.NilObjectID, err
	}
	fmt.Println("validate ===========>", mongoUser.Id_)

	return mongoUser.Id_, nil
}
