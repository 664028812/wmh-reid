package jwt

import (
	"errors"
	"fmt"
	"time"

	"github.com/golang-jwt/jwt"
)

var secretKey = []byte("your-secret-key") // 在实际应用中应该从配置中读取

// Claims 自定义的JWT声明结构
type Claims struct {
	UserID uint   `json:"user_id"`
	Role   string `json:"role"`
	jwt.StandardClaims
}

// GenerateToken 生成JWT token
func GenerateToken(userID uint, role string) (string, error) {
	claims := Claims{
		UserID: userID,
		Role:   role,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(24 * time.Hour).Unix(), // token有效期24小时
			IssuedAt:  time.Now().Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(secretKey)
}

// ValidateToken 验证JWT token
func ValidateToken(tokenString string) (*Claims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("unexpected signing method")
		}
		return secretKey, nil
	})

	if err != nil {
		return nil, err
	}

	if claims, ok := token.Claims.(*Claims); ok && token.Valid {
		return claims, nil
	}

	return nil, errors.New("invalid token")
}

// GenerateRefreshToken 生成刷新令牌
func GenerateRefreshToken(userID uint) (string, error) {
	// 创建刷新令牌的声明
	claims := Claims{
		UserID: userID,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 24 * 7).Unix(), // 7天有效期
			IssuedAt:  time.Now().Unix(),
			Subject:   "refresh_token",
		},
	}

	// 使用专门的刷新令牌密钥
	refreshKey := []byte("your-refresh-secret-key") // 建议从配置文件读取

	// 创建令牌
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	
	// 签名并获取完整的编码后的字符串令牌
	return token.SignedString(refreshKey)
}

// ValidateRefreshToken 验证刷新令牌
func ValidateRefreshToken(tokenString string) (*Claims, error) {
	// 使用与生成时相同的密钥
	refreshKey := []byte("your-refresh-secret-key") // 建议从配置文件读取

	// 解析令牌
	token, err := jwt.ParseWithClaims(tokenString, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		// 验证签名算法
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("意外的签名方法: %v", token.Header["alg"])
		}
		return refreshKey, nil
	})

	if err != nil {
		return nil, fmt.Errorf("解析refresh token失败: %v", err)
	}

	// 验证令牌并转换声明
	if claims, ok := token.Claims.(*Claims); ok && token.Valid {
		// 验证令牌类型
		if claims.Subject != "refresh_token" {
			return nil, fmt.Errorf("无效的refresh token类型")
		}
		return claims, nil
	}

	return nil, fmt.Errorf("无效的refresh token")
}