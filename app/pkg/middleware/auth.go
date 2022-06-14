package middleware

import (
	"strconv"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
	"github.com/mjiee/scaffold-gin/app/lib"
	"github.com/mjiee/scaffold-gin/app/pkg/apperr"
	"github.com/mjiee/scaffold-gin/app/pkg/response"
)

func JwtAuth(iss string, jwtSrv *lib.JwtService, userSrc *lib.UserService) gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.Request.Header.Get("Authorization")
		if authHeader == "" {
			response.Failure(c, apperr.TokenError, "请求头Authorization为空")
			c.Abort()
			return
		}

		// token format validation
		authStr := strings.SplitN(authHeader, " ", 2)
		if len(authStr) != 2 || authStr[0] != "Bearer" {
			response.Failure(c, apperr.TokenError, "请求头Authorization格式错误")
			c.Abort()
			return
		}

		// token parsing
		token, err := jwt.ParseWithClaims(authStr[1], &lib.AppClaims{}, func(token *jwt.Token) (any, error) {
			return []byte(jwtSrv.Conf.Secret), nil
		})
		if err != nil {
			response.Failure(c, apperr.TokenInvalid, err.Error())
			c.Abort()
			return
		}

		claims, ok := token.Claims.(*lib.AppClaims)

		// Issuer validation
		if !ok || claims.Issuer != iss {
			response.Failure(c, apperr.TokenInvalid, "无效token")
			c.Abort()
			return
		}

		// token renew
		if claims.ExpiresAt.Unix()-time.Now().Unix() < int64(30*time.Minute) {
			if !jwtSrv.IsInBlackList(token.Raw) {
				user, err := userSrc.GetUserInfo(claims.ID)
				if err != nil {
					response.Failure(c, apperr.TokenError, err.Error())
					c.Abort()
					return
				} else {
					tokenData, _ := jwtSrv.GenToken(iss, user)
					c.Header("new-token", tokenData.AccessToken)
					c.Header("new-expires-in", strconv.Itoa(tokenData.ExpiresAt))
					_ = jwtSrv.JoinBlackList(token.Raw)
				}
			}
		}

		c.Set("id", claims.ID)
		c.Next()
	}
}
