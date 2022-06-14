package lib

import (
	"context"
	"time"

	"github.com/go-redis/redis/v8"
	"github.com/golang-jwt/jwt/v4"
	"github.com/mjiee/scaffold-gin/app/pkg/conf"
	"github.com/mjiee/scaffold-gin/app/pkg/util"
)

// JwtUser 签名实例需要实现的接口
type JwtUser interface {
	GetUid() string
}

// AppClaims 自定义claims
type AppClaims struct {
	jwt.RegisteredClaims
}

// TokenOutput token输出格式
type TokenOutput struct {
	AccessToken string `json:"access_token"`
	ExpiresAt   int    `json:"expires_At"`
	TokenType   string `json:"token_type"`
}

// JwtAuth 提供jwt相关服务
type JwtService struct {
	Conf  *conf.Jwt
	redis *redis.Client
}

func NewJwtService(cfg *conf.Config, redis *redis.Client) *JwtService {
	return &JwtService{&cfg.Jwt, redis}
}

// GenToken 创建token
func (s *JwtService) GenToken(iss string, user JwtUser) (TokenOutput, error) {
	expiresAt := time.Hour * time.Duration(s.Conf.ExpiresAt)
	claims := AppClaims{
		jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(expiresAt)),
			ID:        user.GetUid(),
			Issuer:    iss,
			//			NotBefore: jwt.NewNumericDate(time.Now().Add(time.Hour)),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenStr, err := token.SignedString([]byte(s.Conf.Secret))
	tokenData := TokenOutput{
		AccessToken: tokenStr,
		ExpiresAt:   int(expiresAt),
		TokenType:   "Bearer",
	}

	return tokenData, err
}

// add token to blacklist
func (s *JwtService) JoinBlackList(tokenStr string) error {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	return s.redis.SetNX(ctx, getBlackListKey(tokenStr), 1, 30*time.Minute).Err()
}

// whether the token is on the blacklist
func (s *JwtService) IsInBlackList(tokenStr string) bool {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	val, err := s.redis.Get(ctx, getBlackListKey(tokenStr)).Result()
	if val == "" || err != nil {
		return false
	}
	return true
}

// get blacklist key
func getBlackListKey(tokenStr string) string {
	return "jwt_black_list:" + util.MD5([]byte(tokenStr))
}
