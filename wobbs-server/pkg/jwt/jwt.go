package jwt

import (
	"errors"
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v4"
)

type CustomClaims struct {
	UserID int64 `json:"user_id"`
	jwt.RegisteredClaims
}

func generateToken(userId int64, expire time.Duration) (string, error) {
	mySigningKey := []byte("AllYourBase")
	// Create the claims
	claims := CustomClaims{
		userId,
		jwt.RegisteredClaims{
			// A usual scenario is to set the expiration time relative to the current time
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(expire)), // 过期时间
			IssuedAt:  jwt.NewNumericDate(time.Now()),             // 签发时间
			NotBefore: jwt.NewNumericDate(time.Now()),             // 生效时间
			Issuer:    "novo",
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(mySigningKey)
	//fmt.Printf("%v %v", tokenString, err)
	return tokenString, err
}

func AccessToken(userId int64) (string, error) {
	return generateToken(userId, 2*time.Hour)
}

func RefreshToken(userId int64) (string, error) {
	return generateToken(userId, 7*time.Hour)
}

func VerifyToken(tokenString string) (*CustomClaims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &CustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte("AllYourBase"), nil
	})

	if err != nil {
		return nil, errors.New("解析token失败")
	}

	if claims, ok := token.Claims.(*CustomClaims); ok && token.Valid {
		//fmt.Printf("%v %v", claims.UserID, claims.RegisteredClaims.Issuer)
		return claims, nil
	}
	return nil, errors.New("token不合法")
}

func main() {
	token, err := AccessToken(23)
	if err != nil {

	}
	fmt.Println(token)
	claims, err := VerifyToken(token)
	fmt.Printf("%v", claims)
}
