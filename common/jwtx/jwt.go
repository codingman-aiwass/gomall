package jwtx

import (
	"errors"
	"fmt"
	"github.com/golang-jwt/jwt"
	"log"
	"time"
)

func GetToken(secretKey string, iat, seconds, uid int64) (string, error) {
	claims := make(jwt.MapClaims)
	claims["exp"] = iat + seconds
	claims["iat"] = iat
	claims["uid"] = uid
	token := jwt.New(jwt.SigningMethodHS256)
	token.Claims = claims
	return token.SignedString([]byte(secretKey))
}

// ValidateToken 验证 JWT 并解析用户信息
func ValidateToken(secretKey string, tokenString string) (uint32, int64, error) {
	// 解析 JWT
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// 验证签名算法
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(secretKey), nil
	})
	if err != nil {
		return 0, 0, err
	}

	// 验证 Claims
	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		// 验证 exp 是否过期
		if exp, ok := claims["exp"].(float64); ok {
			if int64(exp) < jwt.TimeFunc().Unix() {
				return 0, 0, errors.New("token is expired")
			}
		} else {
			return 0, 0, errors.New("invalid exp in token")
		}

		// 解析 uid,过期时间
		uid, ok := claims["uid"].(float64)
		if !ok {
			return 0, 0, errors.New("invalid uid in token")
		}

		exp, ok := claims["exp"].(float64)
		if !ok {
			return 0, 0, errors.New("invalid exp in token")
		}

		now := time.Now().Unix()
		remainingTime := int64(exp) - now
		return uint32(uid), remainingTime, nil
	}

	return 0, 0, errors.New("invalid token")
}

// RenewToken 验证令牌并续期
func RenewToken(refreshSecretKey, accessSecretKey string, seconds int64, accessToken string, refreshToken string) (newAccessToken string, err error) {
	// 如果refresh-token无效，直接返回错误
	_, _, err = ValidateToken(refreshSecretKey, refreshToken)
	if err != nil {
		log.Println("verify refresh token err", err)
		return "", err
	}

	// 解析和验证现有令牌
	token, err := jwt.Parse(accessToken, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(accessSecretKey), nil
	})
	// 如果access-token 有效，不做处理，返回空
	if err != nil {
		v, _ := err.(*jwt.ValidationError)
		// 当错误类型是过期错误，并且refresh token没有过期，创建一个新的access token
		if v.Errors == jwt.ValidationErrorExpired {
			// 提取 uid
			uid, ok := token.Header["uid"].(int64)
			if !ok {
				return "", errors.New("invalid uid in token")
			}
			// 生成新令牌
			return GetToken(accessSecretKey, jwt.TimeFunc().Unix(), seconds, uid)
		}
	}
	return
}
