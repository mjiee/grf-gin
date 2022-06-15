package lib

import (
	"context"
	"errors"
	"strings"
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
	ExpiresAt   int    `json:"expires_at"`
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
		ExpiresAt:   s.Conf.ExpiresAt * 3600,
		TokenType:   "Bearer",
	}

	return tokenData, err
}

// JoinBlackList 将token加入黑名单
func (s *JwtService) JoinBlackList(tokenStr string) error {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	return s.redis.SetNX(ctx, getBlackListKey(tokenStr), 1, 30*time.Minute).Err()
}

// IsInBlackList 查询token是否在黑名单中
func (s *JwtService) IsInBlackList(tokenStr string) bool {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	val, err := s.redis.Get(ctx, getBlackListKey(tokenStr)).Result()
	if val == "" || err != nil {
		return false
	}
	return true
}

// RequestAuth 请求头jwt认证
func (s *JwtService) RequestAuth(iss string, authStr string) (*AppClaims, *jwt.Token, error) {
	if authStr == "" {
		return nil, nil, errors.New("请求头Authorization为空")
	}

	// token格式验证
	authSlice := strings.SplitN(authStr, " ", 2)
	if len(authSlice) != 2 || authSlice[0] != "Bearer" {
		return nil, nil, errors.New("请求头Authorization格式错误")
	}

	// token解析
	token, err := jwt.ParseWithClaims(authSlice[1], &AppClaims{}, func(token *jwt.Token) (any, error) {
		return []byte(s.Conf.Secret), nil
	})

	if err != nil {
		return nil, nil, err
	}

	// 颁发人验证
	claims, ok := token.Claims.(*AppClaims)
	if !ok || claims.Issuer != iss {
		return nil, nil, errors.New("无效token")
	}

	return claims, token, nil
}

// get blacklist key
func getBlackListKey(tokenStr string) string {
	return "jwt_black_list:" + util.MD5([]byte(tokenStr))
}
