package ijwt

import (
	"crmeb_go/define"
	"fmt"
	"github.com/golang-jwt/jwt/v5"
)

type TokenData struct {
	UserId   int64
	ExpireAt int64
}

// ParseToken 解析token
func ParseToken(token string, secretKey string) (*TokenData, error) {
	newToken, err := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(secretKey), nil
	})

	if err != nil {
		return nil, err
	}

	claims, ok := newToken.Claims.(jwt.MapClaims)
	if !ok || !newToken.Valid {
		return nil, fmt.Errorf("invalid token")
	}

	userId, ok := claims[define.AdminUserId].(float64)
	if !ok {
		return nil, fmt.Errorf("invalid user_id type")
	}

	expireAt, ok := claims[define.AdminExpireAt].(float64)
	if !ok {
		return nil, fmt.Errorf("invalid expire_at type")
	}

	return &TokenData{
		UserId:   int64(userId),
		ExpireAt: int64(expireAt),
	}, nil
}

// GenerateToken 生成token
func GenerateToken(userId int64, expiration int64, secretKey string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		define.AdminUserId:   userId,
		define.AdminExpireAt: expiration,
	})

	tokenString, err := token.SignedString([]byte(secretKey))
	if err != nil {
		return "", err
	}
	return tokenString, nil
}
